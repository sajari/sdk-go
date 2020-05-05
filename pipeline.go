package sajari

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes"

	"code.sajari.com/sdk-go/internal/protoutil"

	enginepb "code.sajari.com/protogen-go/sajari/engine/v2"
	pipelinepb "code.sajari.com/protogen-go/sajari/pipeline/v2"
)

// NoDefaultPipelineError is the error type returned when the collection does
// not have a default version set for a given pipeline.
// To resolve errors of this type, the caller should either pass an explicit
// pipeline version along with their pipeline name, or they should set a default
// pipeline version using the API or CLI tools.
type NoDefaultPipelineError struct {
	// Name of the pipeline used in the attempted operation.
	Name string
}

var _ error = (*NoDefaultPipelineError)(nil)

// Error implements error.
func (e *NoDefaultPipelineError) Error() string {
	return fmt.Sprintf("no default version has been set for the pipeline named %q", e.Name)
}

// Pipeline returns a Pipeline for querying a collection.
func (c *Client) Pipeline(name, version string) *Pipeline {
	return &Pipeline{
		name:    name,
		version: version,
		c:       c,
	}
}

// Pipeline is a handler for a named pipeline.
type Pipeline struct {
	name    string
	version string

	c *Client
}

// Search runs a search query defined by a pipeline with the given params and
// session to run in.  Returns the query results and returned params (which could have
// been modified in the pipeline).
func (p *Pipeline) Search(ctx context.Context, params map[string]string, s Session) (*Results, map[string]string, error) {
	pbTracking, err := s.next(params)
	if err != nil {
		return nil, nil, err
	}

	r := &pipelinepb.SearchRequest{
		Pipeline: p.proto(),
		Tracking: pbTracking,
		Values:   protoutil.Struct(params),
	}

	resp, err := pipelinepb.NewQueryClient(p.c.ClientConn).Search(p.c.newContext(ctx), r)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound && strings.HasPrefix(s.Message(), "no default pipeline") {
				err = fmt.Errorf("%w", &NoDefaultPipelineError{Name: p.name})
			}
		}
		return nil, nil, fmt.Errorf("could not run search: %w", err)
	}

	results, err := processResponse(resp.GetQueryResults(), resp.GetTokens()...)
	if err != nil {
		return nil, nil, err
	}

	m, err := protoutil.Map(resp.GetValues())
	if err != nil {
		return nil, nil, err
	}

	return results, m, nil
}

func processResponse(pbResp *pipelinepb.QueryResults, tokens ...*pipelinepb.Token) (*Results, error) {
	pbResults := pbResp.GetResults()
	results := make([]Result, 0, len(pbResults))
	for i, pbr := range pbResults {
		pbValues := pbr.GetValues()
		values := make(map[string]interface{}, len(pbValues))
		for k, v := range pbValues {
			vv, err := protoutil.FromProto(v)
			if err != nil {
				return nil, err
			}
			values[k] = vv
		}

		r := Result{
			Score:      pbr.GetScore(),
			IndexScore: pbr.GetIndexScore(),
			Values:     values,
		}

		if len(tokens) > i {
			switch t := tokens[i].Token.(type) {
			case *pipelinepb.Token_Click_:
				r.Tokens = map[string]interface{}{
					"click": t.Click.GetToken(),
				}

			case *pipelinepb.Token_PosNeg_:
				r.Tokens = map[string]interface{}{
					"pos": t.PosNeg.GetPos(),
					"neg": t.PosNeg.GetNeg(),
				}
			}
		}

		results = append(results, r)
	}

	var d time.Duration
	var err error
	if pbResp.GetLatency() != nil {
		d, err = ptypes.Duration(pbResp.GetLatency())
		if err != nil {
			return nil, err
		}
	}

	resp := &Results{
		Reads:        int(pbResp.GetReads()),
		TotalResults: int(pbResp.GetTotalResults()),
		Latency:      d,
		Results:      results,
	}

	if pbA := pbResp.GetAggregates(); pbA != nil {
		ra, err := processAggregatesResultMap(pbA)
		if err != nil {
			return nil, err
		}
		resp.Aggregates = ra
	}

	if pbA := pbResp.GetAggregateFilters(); pbA != nil {
		ra, err := processAggregatesResultMap(pbA)
		if err != nil {
			return nil, err
		}
		resp.AggregateFilters = ra
	}
	return resp, nil
}

func processAggregateResult(v *enginepb.QueryAggregateResult) (interface{}, error) {
	switch v := v.AggregateResult.(type) {
	case *enginepb.QueryAggregateResult_Count_:
		counts := make(map[string]int, len(v.Count.Counts))
		for ck, cv := range v.Count.Counts {
			counts[ck] = int(cv)
		}
		return CountResult(counts), nil

	case *enginepb.QueryAggregateResult_Buckets_:
		buckets := make(map[string]BucketResult, len(v.Buckets.Buckets))
		for bk, bv := range v.Buckets.Buckets {
			buckets[bk] = BucketResult{
				Name:  bv.Name,
				Count: int(bv.Count),
			}
		}
		return BucketsResult(buckets), nil

	case *enginepb.QueryAggregateResult_Metric_:
		return v.Metric.Value, nil

	case *enginepb.QueryAggregateResult_Date_:
		dates := make(map[string]int, len(v.Date.Dates))
		for ck, cv := range v.Date.Dates {
			dates[ck] = int(cv)
		}
		return DateResult(dates), nil

	case *enginepb.QueryAggregateResult_Analysis_:
		switch vv := v.Analysis.Value.(type) {
		case *enginepb.QueryAggregateResult_Analysis_Coverage:
			return vv.Coverage, nil

		case *enginepb.QueryAggregateResult_Analysis_Cardinality:
			return vv.Cardinality, nil

		case *enginepb.QueryAggregateResult_Analysis_MinLen:
			return vv.MinLen, nil

		case *enginepb.QueryAggregateResult_Analysis_MaxLen:
			return vv.MaxLen, nil

		case *enginepb.QueryAggregateResult_Analysis_AvgLen:
			return vv.AvgLen, nil

		default:
			return nil, fmt.Errorf("unhandled analysis aggregate result: %T", vv)
		}

	default:
		return nil, fmt.Errorf("unhandled aggregate result: %T", v)
	}
}

func processAggregatesResultMap(pbResp map[string]*enginepb.QueryAggregateResult) (map[string]interface{}, error) {
	out := make(map[string]interface{}, len(pbResp))
	for k, v := range pbResp {
		x, err := processAggregateResult(v)
		if err != nil {
			return nil, err
		}
		out[k] = x
	}
	return out, nil
}

// Results is a collection of results from a Search.
type Results struct {
	// Reads is the total number of index values read.
	Reads int

	// TotalResults is the total number of results for the query.
	TotalResults int

	// Time taken to perform the query.
	Latency time.Duration

	// Aggregates computed on the query results (see Aggregate).
	Aggregates map[string]interface{}

	// AggregateFilters computed on query results (see Aggregate).
	AggregateFilters map[string]interface{}

	// Results of the query.
	Results []Result
}

// Result is an individual query result.
type Result struct {
	// Values are field values of records.
	Values map[string]interface{}

	// Tokens contains any tokens associated with this Result.
	Tokens map[string]interface{}

	// Score is the overall score of this Result.
	Score float64

	// IndexScore is the index-matched score of this Result.
	IndexScore float64
}

func (p *Pipeline) proto() *pipelinepb.Identifier {
	return &pipelinepb.Identifier{
		Name:    p.name,
		Version: p.version,
	}
}

// CreateRecord uses a pipeline to add a single record to a collection and
// returns a Key which can be used to retrieve the newly created record.
func (p *Pipeline) CreateRecord(ctx context.Context, values map[string]string, r Record) (*Key, map[string]string, error) {
	pbr, err := r.proto()
	if err != nil {
		return nil, nil, err
	}

	resp, err := pipelinepb.NewStoreClient(p.c.ClientConn).CreateRecord(p.c.newContext(ctx), &pipelinepb.CreateRecordRequest{
		Pipeline: p.proto(),
		Values:   protoutil.Struct(values),
		Record:   pbr,
	})
	if err != nil {
		return nil, nil, err
	}

	k, err := keyFromProto(resp.GetKey())
	if err != nil {
		return nil, nil, err
	}

	m, err := protoutil.Map(resp.GetValues())
	if err != nil {
		return nil, nil, err
	}
	return k, m, nil
}

// ReplaceRecord uses a pipeline to replace a single record in a collection
// represented by the given Key.
func (p *Pipeline) ReplaceRecord(ctx context.Context, values map[string]string, key *Key, r Record) (*Key, map[string]string, error) {
	pbr, err := r.proto()
	if err != nil {
		return nil, nil, err
	}

	pbk, err := key.proto()
	if err != nil {
		return nil, nil, err
	}

	resp, err := pipelinepb.NewStoreClient(p.c.ClientConn).ReplaceRecord(p.c.newContext(ctx), &pipelinepb.ReplaceRecordRequest{
		Pipeline: p.proto(),
		Values:   protoutil.Struct(values),
		Record:   pbr,
		Key:      pbk,
	})
	if err != nil {
		return nil, nil, err
	}

	k, err := keyFromProto(resp.GetKey())
	if err != nil {
		return nil, nil, err
	}

	m, err := protoutil.Map(resp.GetValues())
	if err != nil {
		return nil, nil, err
	}
	return k, m, nil
}

// AggregateResult is an interface implemented by aggregate results.
type AggregateResult interface {
	aggregateResult()
}

// BucketsResult is a type returned from a query performing bucket aggregate.
type BucketsResult map[string]BucketResult

func (BucketsResult) aggregateResult() {}

// BucketResult is bucket information as reported by an aggregate.
type BucketResult struct {
	// Name of the bucket.
	Name string

	// Number of records.
	Count int
}

func (BucketResult) aggregateResult() {}

// CountResult is a type returned from a query which has performed a count aggregate.
type CountResult map[string]int

func (CountResult) aggregateResult() {}

// DateResult is a type returned from a query which has performed a date aggregate.
type DateResult map[string]int

func (DateResult) aggregateResult() {}

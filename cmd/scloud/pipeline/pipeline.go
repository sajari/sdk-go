package pipeline

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	yaml "gopkg.in/yaml.v2"

	"code.sajari.com/api/pipeline"
	"code.sajari.com/api/pipeline/param"
	sajari "code.sajari.com/sdk-go"

	pipelineproto "code.sajari.com/api/pipeline/proto"

	pb "code.sajari.com/protogen-go/sajari/pipeline/v2"
)

var (
	topLevelCommands = []string{"get", "create", "delete", "list", "usage", "step", "steps", "get-default", "set-default", "query", "replace"}
)

// Run executes pipeline control based operations
func Run(client *sajari.Client, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("\nusage: scloud pipeline <%v> [options...]\n\n", strings.Join(topLevelCommands, "|"))
	}

	// Query and replace have different flag sets
	switch args[0] {
	case "query":
		return Query(client, args[1:])

	case "replace":
		return Replace(client, args[1:])
	}

	iflags := flag.NewFlagSet("pipeline", flag.ExitOnError)
	step := iflags.String("step", "", "show step definition for identifier `step`")
	pipelineType := iflags.String("type", "query", "type of pipeline to create (query|record)")
	stepType := iflags.String("stepType", "pre", "type of step to list/get (pre|post)")
	create := iflags.String("create", "", "YAML source path to create pipeline from")
	name := iflags.String("name", "", "pipeline `name`, can be blank if requesting multiple")
	version := iflags.String("version", "", "pipeline `version`")

	if len(args) == 0 {
		defer iflags.Usage()
		return fmt.Errorf("\nusage: scloud pipeline <%v> [options...]\n\n", strings.Join(topLevelCommands, "|"))
	}
	iflags.Parse(args[1:])

	var ty pb.Type
	switch *pipelineType {
	case "query":
		ty = pb.Type_QUERY_PIPELINE
	case "record":
		ty = pb.Type_RECORD_PIPELINE
	default:
		return fmt.Errorf("invalid -type value: %q", *pipelineType)
	}

	var sty pb.StepType
	switch *stepType {
	case "pre":
		sty = pb.StepType_PRE_STEP
	case "post":
		sty = pb.StepType_POST_STEP
	default:
		return fmt.Errorf("invalid -stepType value: %q", *stepType)
	}

	ctx := context.Background()
	ctx = newContext(ctx, client)

	switch args[0] {
	case "get":
		if err := getPipeline(ctx, client, ty, *name, *version); err != nil {
			return fmt.Errorf("Could not get pipeline: %v", err)
		}

	case "create":
		if err := createPipeline(ctx, client, ty, *create); err != nil {
			return fmt.Errorf("Could not create pipeline: %v", err)
		}

	case "delete":
		if err := deletePipeline(ctx, client, ty, *name, *version); err != nil {
			return fmt.Errorf("Could not delete pipeline: %v", err)
		}

	case "list":
		// Since we use non-empty *list to show that we want to run a list,
		// need the sentinal value "-" (which is not a valid pipeline name)
		// to indicate we want the list, but haven't specified a name.

		if err := listPipelines(ctx, client, ty, *name, *version); err != nil {
			return fmt.Errorf("Could not list pipelines: %v", err)
		}

	case "usage":
		if err := pipelineDocumentation(ctx, client, ty, *name, *version); err != nil {
			return fmt.Errorf("Could not fetch documentation for pipeline: %v", err)
		}

	case "get-default":
		if err := getDefaultPipeline(ctx, client, ty, *name, *version); err != nil {
			return fmt.Errorf("Could not get default version: %v", err)
		}

	case "set-default":
		if err := setDefaultPipeline(ctx, client, ty, *name, *version); err != nil {
			return fmt.Errorf("Could not set default version: %v", err)
		}

	case "step":
		sf := pb.NewDocumentationClient(client.ClientConn)
		getStep(ctx, client, sf, ty, *step)

	case "steps":
		sf := pb.NewDocumentationClient(client.ClientConn)
		getSteps(ctx, client, sf, ty, sty)

	default:
		return fmt.Errorf("usage: scloud pipeline <%v> [options...]\n", strings.Join(topLevelCommands, "|"))
	}
	return nil
}

const (
	projectKey    = "project"
	collectionKey = "collection"
)

func newContext(ctx context.Context, client *sajari.Client) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.Pairs(projectKey, client.Project, collectionKey, client.Collection)
	} else {
		md[projectKey] = []string{client.Project}
		md[collectionKey] = []string{client.Collection}
	}
	return metadata.NewOutgoingContext(ctx, md)
}

type tabWriter struct {
	*tabwriter.Writer
}

func (t *tabWriter) reset() {
	t.Init(os.Stdout, 10, 4, 3, ' ', 0)
}

func newTabWriter() *tabWriter {
	return &tabWriter{tabwriter.NewWriter(os.Stdout, 10, 4, 3, ' ', 0)}
}

func (tw *tabWriter) Writef(pattern string, vars ...interface{}) {
	io.WriteString(tw.Writer, fmt.Sprintf(pattern, vars...))
}

func unwrapCodeErr(err error) error {
	if err == nil {
		return nil
	}
	if st, _ := status.FromError(err); st.Code() != codes.Unknown {
		return fmt.Errorf("%v: %v", st.Code(), st.Message())
	}
	return err
}

type stepFetcher interface {
	ListStepTemplates(context.Context, *pb.ListStepTemplatesRequest, ...grpc.CallOption) (*pb.ListStepTemplatesResponse, error)
	GetStepTemplateDocumentation(context.Context, *pb.GetStepTemplateDocumentationRequest, ...grpc.CallOption) (*pb.GetStepTemplateDocumentationResponse, error)
}

func getSteps(ctx context.Context, c *sajari.Client, sf stepFetcher, ty pb.Type, sty pb.StepType) {

	var pageToken string
	var steps []*pb.Step
	for {
		resp, err := sf.ListStepTemplates(ctx, &pb.ListStepTemplatesRequest{
			Type:      ty,
			StepType:  sty,
			PageToken: pageToken,
		})
		if err != nil {
			log.Fatalf("Could not list %v pipeline steps: %v", ty.String(), unwrapCodeErr(err))
		}
		steps = append(steps, resp.GetStepTemplates()...)
		pageToken = resp.GetNextPageToken()
		if pageToken == "" {
			break
		}
	}

	tw := newTabWriter()
	tw.Writef("ID\tTITLE\tDESCRIPTION\n")

	for _, s := range steps {
		tw.Writef("%v\t%v\t%v\n", s.GetIdentifier(), s.GetTitle(), s.GetDescription())
	}
	if err := tw.Flush(); err != nil {
		log.Fatalf("Could not flush writer: %v", err)
	}
}

func getStep(ctx context.Context, c *sajari.Client, sf stepFetcher, ty pb.Type, id string) {

	resp, err := sf.GetStepTemplateDocumentation(ctx, &pb.GetStepTemplateDocumentationRequest{
		Type:       ty,
		Identifier: id,
	})
	if err != nil {
		log.Fatalf("Could not get %v pipeine step %q: %v", ty.String(), id, unwrapCodeErr(err))
	}

	st := resp.GetStepTemplate()

	tw := newTabWriter()
	tw.Writef("Type:\t%v\n", st.GetType())
	tw.Writef("Step Types:\t%v\n", st.GetStepTypes())
	tw.Writef("Identifier:\t%v\n", st.GetIdentifier())
	tw.Writef("Title:\t%v\n", st.GetTitle())
	tw.Writef("Description:\t%v\n", st.GetDescription())
	if err := tw.Flush(); err != nil {
		log.Fatalf("Could not flush writer: %v", err)
	}

	if len(st.GetInputs()) > 0 {
		fmt.Println("\nINPUTS")
		writeParamTable(st.GetInputs())
	}

	if len(st.GetOutputs()) > 0 {
		fmt.Println("\nOUTPUTS")
		writeParamTable(st.GetOutputs())
	}

	if len(st.GetConstants()) > 0 {
		fmt.Println("\nCONSTANTS")
		writeConstTable(st.GetConstants())
	}

	if err := tw.Flush(); err != nil {
		log.Fatalf("Could not flush writer: %v", err)
	}
}

func writeParamTable(ps []*pb.Parameter) {
	tw := newTabWriter()
	tw.Writef("ID\tNAME\tTYPE\tDEFAULT VALUE\tDESCRIPTION\n")
	for _, p := range ps {
		tw.Writef("%v\t%v\t%v\t%q\t%v\n", p.GetIdentifier(), p.GetName(), p.GetType(), p.GetDefaultValue(), p.GetDescription())
	}
	if err := tw.Flush(); err != nil {
		log.Fatalf("Could not flush writer: %v", err)
	}
}

func writeConstTable(ps []*pb.Constant) {
	tw := newTabWriter()
	tw.Writef("NAME\tTYPE\tVALUE\tDESCRIPTION\n")
	for _, p := range ps {
		tw.Writef("%v\t%v\t%q\t%v\n", p.GetName(), p.GetType(), p.GetValue(), p.GetDescription())
	}
	if err := tw.Flush(); err != nil {
		log.Fatalf("Could not flush writer: %v", err)
	}
}

func deletePipeline(ctx context.Context, c *sajari.Client, ty pb.Type, name, version string) error {
	_, err := pb.NewPipelineAdminClient(c.ClientConn).DeletePipeline(ctx, &pb.DeletePipelineRequest{
		Type: ty,
		Pipeline: &pb.Identifier{
			Name:    name,
			Version: version,
		},
	})
	return unwrapCodeErr(err)
}

func listPipelines(ctx context.Context, c *sajari.Client, ty pb.Type, name, version string) error {
	var pageToken string
	var ps []*pb.Pipeline
	for {
		resp, err := pb.NewPipelineAdminClient(c.ClientConn).ListPipelines(ctx, &pb.ListPipelinesRequest{
			Pipeline:  identifier(ty, name, version),
			PageToken: pageToken,
		})
		if err != nil {
			return err
		}

		ps = append(ps, resp.GetPipelines()...)

		pageToken = resp.GetNextPageToken()
		if pageToken == "" {
			break
		}
	}

	tw := newTabWriter()
	tw.Writef("TYPE\tNAME\tVERSION\tCREATED\n")
	for _, p := range ps {
		x := p.GetPipeline()
		t, err := ptypes.Timestamp(p.GetCreateTime())
		if err != nil {
			return fmt.Errorf("could not parse timestamp: %v", err)
		}
		tw.Writef(
			"%v\t%v\t%v\t%v\t\n",
			x.GetType(),
			x.GetIdentifier().GetName(),
			x.GetIdentifier().GetVersion(),
			humanTime(t),
		)
	}
	if err := tw.Flush(); err != nil {
		log.Fatalf("Could not flush writer: %v", err)
	}
	return nil
}

type humanTime time.Time

func (t humanTime) String() string {
	return fmt.Sprintf(
		"%v (%v)",
		time.Time(t).Local().Format("2006-01-02 15:04:05 -0700 MST"),
		duration(time.Since(time.Time(t))),
	)
}

// duration is a wrapper for creating more human-friendly duration
// reps.
type duration time.Duration

func (d duration) String() string {
	switch x := time.Duration(d); {
	case x > 24*time.Hour:
		return fmt.Sprintf("%d days ago", x/(24*time.Hour))
	case x > 2*time.Hour:
		return fmt.Sprintf("%d hours ago", x/time.Hour)
	case x > 2*time.Minute:
		return fmt.Sprintf("%d mins ago", x/time.Minute)
	default:
		return fmt.Sprintf("%d secs ago", x/time.Second)
	}
}

func identifier(ty pb.Type, name, version string) *pb.TypeIdentifier {
	return &pb.TypeIdentifier{
		Type: ty,
		Identifier: &pb.Identifier{
			Name:    name,
			Version: version,
		},
	}
}

func setDefaultPipeline(ctx context.Context, c *sajari.Client, ty pb.Type, name, version string) error {
	_, err := pb.NewPipelineAdminClient(c.ClientConn).SetDefaultPipeline(ctx, &pb.SetDefaultPipelineRequest{
		Pipeline: identifier(ty, name, version),
	})
	return unwrapCodeErr(err)
}

func getDefaultPipeline(ctx context.Context, c *sajari.Client, ty pb.Type, name, version string) error {
	resp, err := pb.NewPipelineAdminClient(c.ClientConn).GetDefaultPipeline(ctx, &pb.GetDefaultPipelineRequest{
		Pipeline: identifier(ty, name, version),
	})
	if err != nil {
		return unwrapCodeErr(err)
	}

	t, err := ptypes.Timestamp(resp.GetCreateTime())
	if err != nil {
		log.Fatalf("Could not parse timestamp value: %v", err)
	}

	tw := newTabWriter()
	tw.Writef("Type:\t%v\n", resp.GetPipeline().GetType())
	tw.Writef("Name:\t%q\n", resp.GetPipeline().GetIdentifier().GetName())
	tw.Writef("Version:\t%q\n", resp.GetPipeline().GetIdentifier().GetVersion())
	tw.Writef("Created:\t%v\n\n", humanTime(t))

	if err := tw.Flush(); err != nil {
		log.Fatalf("Could not flush writer: %v", err)
	}
	return nil
}

type getParams interface {
	GetInputs() []*pb.Parameter
	GetOutputs() []*pb.Parameter
}

func writeParams(gp getParams) {
	var inputs bool
	if ps := gp.GetInputs(); len(ps) > 0 {
		fmt.Println("INPUTS")
		fmt.Println("------")
		writeUsageParamTable(ps)
		inputs = true
	}
	if ps := gp.GetOutputs(); len(ps) > 0 {
		if inputs {
			fmt.Println()
		}
		fmt.Println("OUTPUTS")
		fmt.Println("-------")
		writeUsageParamTable(ps)
	}
}

func writeSteps(name string, gs []*pb.Step) {
	if len(gs) > 0 {
		fmt.Println(name)
		fmt.Println(strings.Repeat("-", len(name)))
		writeStepInfo(gs)
	}
}

func pipelineDocumentation(ctx context.Context, c *sajari.Client, ty pb.Type, name, version string) error {
	resp, err := pb.NewDocumentationClient(c.ClientConn).GetPipelineDocumentation(ctx, &pb.GetPipelineDocumentationRequest{
		Pipeline: identifier(ty, name, version),
	})
	if err != nil {
		return unwrapCodeErr(err)
	}

	t, err := ptypes.Timestamp(resp.GetCreateTime())
	if err != nil {
		log.Fatalf("Could not parse timestamp value: %v", err)
	}

	tw := newTabWriter()
	tw.Writef("Type:\t%v\n", resp.GetPipeline().GetType())
	tw.Writef("Name:\t%q\n", resp.GetPipeline().GetIdentifier().GetName())
	tw.Writef("Version:\t%q\n", resp.GetPipeline().GetIdentifier().GetVersion())
	tw.Writef("Created:\t%v\n\n", humanTime(t))

	tw.Writef("Title:\t%v\n", resp.GetTitle())
	tw.Writef("Description:\t%v\n\n", resp.GetDescription())

	if err := tw.Flush(); err != nil {
		log.Fatalf("Could not flush writer: %v", err)
	}

	var preSteps, postSteps []*pb.Step
	for _, scs := range resp.GetSteps() {
		switch st := scs.GetStepType(); st {
		case pb.StepType_PRE_STEP:
			preSteps = scs.GetSteps()

		case pb.StepType_POST_STEP:
			postSteps = scs.GetSteps()

		default:
			return fmt.Errorf("unknown step type: %v", st)
		}
	}

	writeParams(resp)
	fmt.Println()
	writeSteps("PRE-STEPS", preSteps)
	fmt.Println()
	writeSteps("POST-STEPS", postSteps)
	return nil
}

func writeStepInfo(ps []*pb.Step) {
	tw := newTabWriter()
	tw.Writef("IDENTIFIER\tNAME\tDESCRIPTION\tCONDITION\n")
	for _, p := range ps {
		tw.Writef("%v\t%v\t%v\t%q\n", p.GetIdentifier(), p.GetTitle(), p.GetDescription(), p.GetCondition())
	}
	if err := tw.Flush(); err != nil {
		log.Fatalf("Could not flush writer: %v", err)
	}
}

func writeUsageParamTable(ps []*pb.Parameter) {
	tw := newTabWriter()
	tw.Writef("NAME\tTYPE\tDEFAULT VALUE\tDESCRIPTION\n")
	for _, p := range ps {
		tw.Writef("%v\t%v\t%q\t%v\n", p.GetName(), p.GetType(), p.GetDefaultValue(), p.GetDescription())
	}
	if err := tw.Flush(); err != nil {
		log.Fatalf("Could not flush writer: %v", err)
	}
}

func createPipeline(ctx context.Context, c *sajari.Client, ty pb.Type, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	var cp createPipelineFile
	if err := yaml.Unmarshal(b, &cp); err != nil {
		return err
	}

	r, err := cp.proto()
	if err != nil {
		return err
	}
	r.Pipeline.Pipeline.Type = ty

	if _, err := pb.NewPipelineAdminClient(c.ClientConn).CreatePipeline(ctx, r); err != nil {
		return unwrapCodeErr(err)
	}
	return nil
}

func getPipeline(ctx context.Context, c *sajari.Client, ty pb.Type, name, version string) (err error) {
	resp, err := pb.NewPipelineAdminClient(c.ClientConn).GetPipeline(ctx, &pb.GetPipelineRequest{
		Pipeline: identifier(ty, name, version),
	})
	if err != nil {
		return unwrapCodeErr(err)
	}

	creator, err := pipelineproto.ParseCreatePipeline(resp.GetPipeline())
	if err != nil {
		return err
	}

	cpf, err := FromCreator(creator)
	if err != nil {
		return err
	}

	var w io.Writer = os.Stdout
	enc := yaml.NewEncoder(w)
	defer func() {
		if err1 := enc.Close(); err1 != nil && err == nil {
			err = err1
		}
	}()

	if err := enc.Encode(cpf); err != nil {
		return err
	}
	return nil
}

func FromParamOpt(p param.ParamOpt) (ParameterConfig, error) {
	switch x := p.(type) {
	case param.SetName:
		return ParameterConfig{
			SetName: string(x),
		}, nil

	case param.SetDescription:
		return ParameterConfig{
			SetDescription: string(x),
		}, nil

	case param.SetDefaultValue:
		return ParameterConfig{
			SetDefaultValue: string(x),
		}, nil

	default:
		return ParameterConfig{}, fmt.Errorf("unknown param option: %T", x)
	}
}

type ParameterConfig struct {
	SetName         string `yaml:"name,omitempty"`
	SetDescription  string `yaml:"description,omitempty"`
	SetDefaultValue string `yaml:"defaultValue,omitempty"`
}

func (p ParameterConfig) paramOpt() (param.ParamOpt, error) {
	if p.SetName != "" {
		return param.SetName(p.SetName), nil
	}

	if p.SetDescription != "" {
		return param.SetDescription(p.SetDescription), nil
	}

	if p.SetDefaultValue != "" {
		return param.SetDefaultValue(p.SetDefaultValue), nil
	}
	return nil, fmt.Errorf("empty param opt")
}

func FromConstOpt(c param.ConstOpt) (ConstantConfig, error) {
	if x, ok := c.(param.SetValue); ok {
		return ConstantConfig{
			SetValue: string(x),
		}, nil
	}
	return ConstantConfig{}, fmt.Errorf("unknown constant option: %T", c)
}

type ConstantConfig struct {
	SetValue string `yaml:"value"`
}

func (c ConstantConfig) constOpt() param.ConstOpt {
	// Allow empty value, as this could be required to override
	// a default.
	return param.SetValue(c.SetValue)
}

type StepConfig struct {
	Identifier  string `yaml:"id"`
	Title       string `yaml:"title,omitempty"`
	Description string `yaml:"description,omitempty"`

	ParameterConfigs map[string][]ParameterConfig `yaml:"params,omitempty"`
	ConstantConfigs  map[string][]ConstantConfig  `yaml:"consts,omitempty"`

	Condition string `yaml:"condition,omitempty"`
}

func FromStepCreator(creator pipeline.StepCreator) (StepConfig, error) {
	ccs := make(map[string][]ConstantConfig, len(creator.ConstOpts))
	for id, co := range creator.ConstOpts {
		for _, x := range co {
			cc, err := FromConstOpt(x)
			if err != nil {
				return StepConfig{}, err
			}
			ccs[id] = append(ccs[id], cc)
		}
	}

	pcs := make(map[string][]ParameterConfig, len(creator.ParamOpts))
	for id, po := range creator.ParamOpts {
		for _, x := range po {
			pc, err := FromParamOpt(x)
			if err != nil {
				return StepConfig{}, err
			}
			pcs[id] = append(pcs[id], pc)
		}
	}

	return StepConfig{
		Identifier:       creator.ID,
		Title:            creator.Title,
		Description:      creator.Description,
		Condition:        creator.Condition,
		ConstantConfigs:  ccs,
		ParameterConfigs: pcs,
	}, nil
}

func (c StepConfig) stepCreator() (pipeline.StepCreator, error) {
	paramOpts := make(map[string][]param.ParamOpt, len(c.ParameterConfigs))
	for id, pos := range c.ParameterConfigs {
		x := make([]param.ParamOpt, 0, len(pos))
		for _, po := range pos {
			z, err := po.paramOpt()
			if err != nil {
				return pipeline.StepCreator{}, err
			}
			x = append(x, z)
		}
		paramOpts[id] = x
	}

	constOpts := make(map[string][]param.ConstOpt, len(c.ConstantConfigs))
	for id, cos := range c.ConstantConfigs {
		x := make([]param.ConstOpt, 0, len(cos))
		for _, co := range cos {
			x = append(x, co.constOpt())
		}
		constOpts[id] = x
	}

	return pipeline.StepCreator{
		ID:          c.Identifier,
		Title:       c.Title,
		Description: c.Description,
		ParamOpts:   paramOpts,
		ConstOpts:   constOpts,
		Condition:   c.Condition,
	}, nil
}

func FromStepCreators(scs []pipeline.StepCreator) ([]StepConfig, error) {
	out := make([]StepConfig, 0, len(scs))
	for _, sc := range scs {
		c, err := FromStepCreator(sc)
		if err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, nil
}

func FromCreator(c *pipeline.Creator) (createPipelineFile, error) {
	preSteps, err := FromStepCreators(c.PreSteps)
	if err != nil {
		return createPipelineFile{}, err
	}

	postSteps, err := FromStepCreators(c.PostSteps)
	if err != nil {
		return createPipelineFile{}, err
	}

	return createPipelineFile{
		Name:        c.Name(),
		Version:     c.Version(),
		Title:       c.Title(),
		Description: c.Description(),
		PreSteps:    preSteps,
		PostSteps:   postSteps,
	}, nil
}

type createPipelineFile struct {
	Name        string       `yaml:"name,omitempty"`
	Version     string       `yaml:"version,omitempty"`
	Title       string       `yaml:"title,omitempty"`
	Description string       `yaml:"description,omitempty"`
	PreSteps    []StepConfig `yaml:"pre-steps,omitempty"`
	PostSteps   []StepConfig `yaml:"post-steps,omitempty"`
}

// newInfo creates a new Info instance.
func newInfo(name, version, title, description string) *info {
	return &info{
		name:        name,
		version:     version,
		title:       title,
		description: description,
	}
}

// info implements pipeline.Info and describes a resource.
type info struct {
	name, version, title, description string
}

// Name is the name of the resource.
func (d *info) Name() string { return d.name }

// Version is the version of the resource.
func (d *info) Version() string { return d.version }

// Title is a human-readable description of the pipeline.
func (d *info) Title() string { return d.title }

// Description is a human-readable description of the pipeline.
func (d *info) Description() string { return d.description }

func (c *createPipelineFile) proto() (*pb.CreatePipelineRequest, error) {
	steps := make([]pipeline.StepCreator, 0, len(c.PreSteps))
	for _, s := range c.PreSteps {
		step, err := s.stepCreator()
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}

	postSteps := make([]pipeline.StepCreator, 0, len(c.PostSteps))
	for _, s := range c.PostSteps {
		step, err := s.stepCreator()
		if err != nil {
			return nil, err
		}
		postSteps = append(postSteps, step)
	}

	creator := &pipeline.Creator{
		Info:      newInfo(c.Name, c.Version, c.Title, c.Description),
		PreSteps:  steps,
		PostSteps: postSteps,
	}

	pbPipeline, err := pipelineproto.FromCreator(creator)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePipelineRequest{
		Pipeline: pbPipeline,
	}, nil
}

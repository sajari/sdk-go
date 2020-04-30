package sajari

import (
	"context"
	"errors"
	"fmt"

	"code.sajari.com/sdk-go/internal/protoutil"

	enginepb "code.sajari.com/protogen-go/sajari/engine/v2"
)

// ErrNoSuchRecord is returned when a requested record cannot be found.
var ErrNoSuchRecord = errors.New("sajari: no such record")

// Record is a set of field-value pairs representing a record in a collection.
type Record map[string]interface{}

func (r Record) proto() (*enginepb.Record, error) {
	values, err := protoutil.Values(r)
	if err != nil {
		return nil, err
	}
	return &enginepb.Record{
		Values: values,
	}, nil
}

// Key is a unique identifier record.
type Key struct {
	field string
	value interface{}
}

// NewKey creates a new Key with a field and value.  Field must be marked as unique in
// the collection schema.
func NewKey(field string, value interface{}) *Key {
	return &Key{
		field: field,
		value: value,
	}
}

// String implements Stringer.
func (k *Key) String() string {
	if k == nil {
		return ""
	}
	return fmt.Sprintf("Key{Field: %q, Value: %q}", k.field, k.value)
}

type keys []*Key

func (ks keys) proto() ([]*enginepb.Key, error) {
	out := make([]*enginepb.Key, 0, len(ks))
	for _, k := range ks {
		pbk, err := k.proto()
		if err != nil {
			return nil, err
		}
		out = append(out, pbk)
	}
	return out, nil
}

func (k *Key) proto() (*enginepb.Key, error) {
	if k == nil {
		return nil, fmt.Errorf("empty key")
	}
	vv, err := protoutil.Single(k.value)
	if err != nil {
		return nil, fmt.Errorf("error marshalling key value: %v", err)
	}
	return &enginepb.Key{
		Field: k.field,
		Value: vv,
	}, nil
}

func keyFromProto(k *enginepb.Key) (*Key, error) {
	if k == nil {
		return nil, nil
	}

	if k.Field == "" && k.Value == nil {
		return nil, nil
	}
	val, err := protoutil.FromProto(k.Value)
	if err != nil {
		return nil, err
	}
	return NewKey(k.Field, val), nil
}

type pbKeys []*enginepb.Key

func (pbks pbKeys) keys() ([]*Key, error) {
	out := make([]*Key, 0, len(pbks))
	for _, pbKey := range pbks {
		key, err := keyFromProto(pbKey)
		if err != nil {
			return nil, err
		}
		out = append(out, key)
	}
	return out, nil
}

func recordFromProto(pbr *enginepb.Record) (Record, error) {
	d := make(Record)
	for k, v := range pbr.Values {
		vv, err := protoutil.FromProto(v)
		if err != nil {
			return nil, err
		}
		d[k] = vv
	}
	return d, nil
}

func (c *Client) MutateRecord(ctx context.Context, k *Key, fms ...RecordMutation) error {
	pbk, err := k.proto()
	if err != nil {
		return err
	}

	pbfms, err := recordMutations(fms).proto()
	if err != nil {
		return err
	}

	_, err = enginepb.NewStoreClient(c.ClientConn).MutateRecord(c.newContext(ctx), &enginepb.MutateRecordRequest{
		Key:            pbk,
		FieldMutations: pbfms,
	})
	return err
}

type recordMutations []RecordMutation

func (rms recordMutations) proto() ([]*enginepb.MutateRecordRequest_FieldMutation, error) {
	out := make([]*enginepb.MutateRecordRequest_FieldMutation, 0, len(rms))
	for _, rm := range rms {
		rmpb, err := rm.proto()
		if err != nil {
			return nil, err
		}
		out = append(out, rmpb)
	}
	return out, nil
}

// DeleteRecord removes a record identified by the key k.  Returns non-nil error if there was
// a communication problem, but fails silently if any key doesn't have a corresponding record.
func (c *Client) DeleteRecord(ctx context.Context, k *Key) error {
	pbk, err := k.proto()
	if err != nil {
		return err
	}

	_, err = enginepb.NewStoreClient(c.ClientConn).DeleteRecord(c.newContext(ctx), &enginepb.DeleteRecordRequest{
		Key: pbk,
	})
	return err
}

// GetRecord retrives the record identified by the key k.
func (c *Client) GetRecord(ctx context.Context, k *Key) (Record, error) {
	pbk, err := k.proto()
	if err != nil {
		return nil, err
	}

	resp, err := enginepb.NewStoreClient(c.ClientConn).GetRecord(c.newContext(ctx), &enginepb.GetRecordRequest{
		Key: pbk,
	})
	if err != nil {
		return nil, err
	}
	return recordFromProto(resp.GetRecord())
}

// Keys returns an iterator which will retrieve the given key field value
// for each record in the collection. If changes to the collection are made
// whilst iterating, the iterator may become invalid or return keys already
// visited.
func (c *Client) Keys(ctx context.Context, field string) *KeyIterator {
	return &KeyIterator{
		ctx:   ctx,
		c:     c,
		field: field,
	}
}

// ErrDone is returned when the iteration is complete.
var ErrDone = errors.New("done")

// KeyIterator iterates through a list of keys.
type KeyIterator struct {
	ctx     context.Context
	c       *Client
	field   string
	token   string
	keys    []*Key
	end     bool
	lastErr error
}

// Next returns the next key in the iteration, or nil, ErrDone if there
// are no more keys remaining.
func (it *KeyIterator) Next() (*Key, error) {
	if it.lastErr != nil {
		return nil, it.lastErr
	}
	if len(it.keys) == 0 && it.end {
		return nil, ErrDone
	}

	if len(it.keys) == 0 {
		if it.keys, it.token, it.lastErr = it.fetch(it.ctx); it.lastErr != nil {
			return nil, it.lastErr
		}
		if it.token == "" {
			it.end = true
		}

		if len(it.keys) == 0 {
			return nil, ErrDone
		}
	}

	k := it.keys[0]
	it.keys = it.keys[1:]
	return k, nil
}

func (it *KeyIterator) fetch(ctx context.Context) (ks []*Key, token string, err error) {
	resp, err := enginepb.NewStoreClient(it.c.ClientConn).ListKeys(it.c.newContext(ctx), &enginepb.ListKeysRequest{
		Field:     it.field,
		PageSize:  0,
		PageToken: it.token,
	})
	if err != nil {
		return nil, "", err
	}

	ks, err = pbKeys(resp.GetKeys()).keys()
	if err != nil {
		return nil, "", err
	}
	return ks, resp.GetNextPageToken(), nil
}

// SetFields is a convenience method for creating field mutations
// to set a map of values.
func SetFields(m map[string]interface{}) []RecordMutation {
	out := make([]RecordMutation, 0, len(m))
	for k, v := range m {
		out = append(out, SetFieldValue(k, v))
	}
	return out
}

// RecordMutation is an interface satisfied by all record mutations defined
// in this package.
type RecordMutation interface {
	proto() (*enginepb.MutateRecordRequest_FieldMutation, error)
}

type setField struct {
	field string
	value interface{}
}

func (s setField) proto() (*enginepb.MutateRecordRequest_FieldMutation, error) {
	v, err := protoutil.Value(s.value)
	if err != nil {
		return nil, err
	}

	return &enginepb.MutateRecordRequest_FieldMutation{
		Field: s.field,
		Mutation: &enginepb.MutateRecordRequest_FieldMutation_Set{
			Set: v,
		},
	}, nil
}

// SetFieldValue is a RecordMutation which sets field to value.  If value is nil
// then this unsets field.
func SetFieldValue(field string, value interface{}) RecordMutation {
	return setField{field, value}
}

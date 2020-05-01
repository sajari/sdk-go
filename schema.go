package sajari

import (
	"context"
	"fmt"

	pb "code.sajari.com/protogen-go/sajari/engine/v2"
)

// Schema returns the schema (list of fields) for the collection.
func (c *Client) Schema() *Schema {
	return &Schema{
		c: c,
	}
}

// Schema provides methods for managing collection schemas.  Use Client.Schema to create
// one for a collection.
type Schema struct {
	c *Client
}

// Fields returns an iterator which retrieves all the fields in the collection.
func (s *Schema) Fields(ctx context.Context) *FieldIterator {
	return &FieldIterator{
		ctx: ctx,
		c:   s.c,
	}
}

// FieldIterator iterates through a list of fields.
type FieldIterator struct {
	ctx     context.Context
	c       *Client
	token   string
	fields  []Field
	end     bool
	lastErr error
}

// Next returns the next field in the iteration. If there are no more fields
// remaining then an error wrapping ErrDone is returned.
func (it *FieldIterator) Next() (Field, error) {
	if it.lastErr != nil {
		return Field{}, it.lastErr
	}
	if len(it.fields) == 0 && it.end {
		return Field{}, fmt.Errorf("%w", ErrDone)
	}

	if len(it.fields) == 0 {
		if it.fields, it.token, it.lastErr = it.fetch(it.ctx); it.lastErr != nil {
			return Field{}, it.lastErr
		}
		if it.token == "" {
			it.end = true
		}
	}

	f := it.fields[0]
	it.fields = it.fields[1:]
	return f, nil
}

func (it *FieldIterator) fetch(ctx context.Context) ([]Field, string, error) {
	resp, err := pb.NewSchemaClient(it.c.ClientConn).ListFields(it.c.newContext(ctx),
		&pb.ListFieldsRequest{
			PageToken: it.token,
		})
	if err != nil {
		return nil, "", err
	}

	fs := make([]Field, 0, len(resp.GetFields()))
	for _, pbField := range resp.GetFields() {
		f, err := fieldFromProto(pbField)
		if err != nil {
			return nil, "", err
		}
		fs = append(fs, f)
	}
	return fs, resp.GetNextPageToken(), nil
}

func fieldFromProto(f *pb.Field) (Field, error) {
	t, err := typeFromProto(f.GetType())
	if err != nil {
		return Field{}, err
	}

	m, err := modeFromProto(f.GetMode())
	if err != nil {
		return Field{}, err
	}

	pbIndexes := f.GetIndexes()
	indexes := make([]FieldIndex, 0, len(pbIndexes))
	for _, pbIndex := range pbIndexes {
		indexes = append(indexes, FieldIndex{
			Spec:        pbIndex.GetSpec(),
			Description: pbIndex.GetDescription(),
		})
	}

	return Field{
		Name:        f.GetName(),
		Description: f.GetDescription(),
		Type:        t,
		Mode:        m,
		Repeated:    f.GetRepeated(),
		Indexes:     indexes,
	}, nil
}

// Field represents a meta field which can be assigned in a collection record.
type Field struct {
	// Name is the name used to identify the field.
	Name string

	// Description is a description of the field.
	Description string

	// Type defines the type of the field.
	Type FieldType

	// Mode of the field.
	Mode FieldMode

	// Repeated indicates that this field can hold a list of values.
	Repeated bool

	// Indexes is a list of the field's indexes.
	Indexes []FieldIndex
}

// Index returns the index matching the given identifier/specification.
// If no such index exists then it will return ({}, false).
func (f Field) Index(spec string) (FieldIndex, bool) {
	for _, idx := range f.Indexes {
		if idx.Spec == spec {
			return idx, true
		}
	}
	return FieldIndex{}, false
}

// FieldIndex is a field index.
type FieldIndex struct {
	// Spec is the identifier/specification for the creation of the index.
	Spec string
	// Description is a description of the index.
	Description string
}

func (f FieldIndex) proto() *pb.FieldIndex {
	return &pb.FieldIndex{
		Spec:        f.Spec,
		Description: f.Description,
	}
}

func (f Field) proto() (*pb.Field, error) {
	t, err := f.Type.proto()
	if err != nil {
		return nil, err
	}

	m, err := f.Mode.proto()
	if err != nil {
		return nil, err
	}

	return &pb.Field{
		Name:        f.Name,
		Description: f.Description,
		Type:        t,
		Mode:        m,
		Repeated:    f.Repeated,
	}, nil
}

func typeFromProto(t pb.Field_Type) (FieldType, error) {
	switch t {
	case pb.Field_STRING:
		return TypeString, nil

	case pb.Field_INTEGER:
		return TypeInteger, nil

	case pb.Field_FLOAT:
		return TypeFloat, nil

	case pb.Field_DOUBLE:
		return TypeDouble, nil

	case pb.Field_BOOLEAN:
		return TypeBoolean, nil

	case pb.Field_TIMESTAMP:
		return TypeTimestamp, nil

	default:
		return TypeString, fmt.Errorf("unknown type: '%v'", t)
	}
}

func modeFromProto(m pb.Field_Mode) (FieldMode, error) {
	switch m {
	case pb.Field_NULLABLE:
		return ModeNullable, nil

	case pb.Field_REQUIRED:
		return ModeRequired, nil

	case pb.Field_UNIQUE:
		return ModeUnique, nil

	default:
		return ModeNullable, fmt.Errorf("unknown mode: '%v'", m)
	}
}

// FieldMode defines field modes.
type FieldMode string

// Enumeration of field modes.
const (
	ModeNullable FieldMode = "NULLABLE" // Don't require a value.
	ModeRequired FieldMode = "REQUIRED" // Field value must be set.
	ModeUnique   FieldMode = "UNIQUE"   // Field value must be unique (and hence also set).
)

func (m FieldMode) proto() (pb.Field_Mode, error) {
	switch m {
	case ModeNullable:
		return pb.Field_NULLABLE, nil
	case ModeRequired:
		return pb.Field_REQUIRED, nil
	case ModeUnique:
		return pb.Field_UNIQUE, nil
	default:
		return pb.Field_NULLABLE, fmt.Errorf("unknown mode: %q", string(m))
	}
}

// FieldType defines field data types.
type FieldType string

// Enumeration of field types.
const (
	TypeString    FieldType = "STRING"
	TypeInteger   FieldType = "INTEGER"
	TypeFloat     FieldType = "FLOAT"
	TypeDouble    FieldType = "DOUBLE"
	TypeBoolean   FieldType = "BOOLEAN"
	TypeTimestamp FieldType = "TIMESTAMP"
)

func (t FieldType) proto() (pb.Field_Type, error) {
	switch t {
	case TypeString:
		return pb.Field_STRING, nil

	case TypeInteger:
		return pb.Field_INTEGER, nil

	case TypeFloat:
		return pb.Field_FLOAT, nil

	case TypeDouble:
		return pb.Field_DOUBLE, nil

	case TypeBoolean:
		return pb.Field_BOOLEAN, nil

	case TypeTimestamp:
		return pb.Field_TIMESTAMP, nil
	}
	return pb.Field_STRING, fmt.Errorf("unknown type: '%v'", string(t))
}

// CreateField creates a new field in the schema.
func (s *Schema) CreateField(ctx context.Context, f Field) error {
	pbf, err := f.proto()
	if err != nil {
		return err
	}

	_, err = pb.NewSchemaClient(s.c.ClientConn).CreateField(s.c.newContext(ctx), &pb.CreateFieldRequest{
		Field: pbf,
	})
	return err
}

// MutateField mutates the field identified by name.
func (s *Schema) MutateField(ctx context.Context, name string, m FieldMutation) error {
	pbm, err := m.proto()
	if err != nil {
		return err
	}

	_, err = pb.NewSchemaClient(s.c.ClientConn).MutateField(ctx, &pb.MutateFieldRequest{
		Name:     name,
		Mutation: pbm,
	})
	return err
}

type fieldMutations []FieldMutation

func (ms fieldMutations) proto() ([]*pb.MutateFieldRequest_Mutation, error) {
	out := make([]*pb.MutateFieldRequest_Mutation, 0, len(ms))
	for _, m := range ms {
		x, err := m.proto()
		if err != nil {
			return nil, err
		}
		out = append(out, x)
	}
	return out, nil
}

// FieldNameMutation creates a schema field mutation which changes the name of a field.
func FieldNameMutation(name string) FieldMutation {
	return fieldNameMutation(name)
}

type fieldNameMutation string

func (n fieldNameMutation) proto() (*pb.MutateFieldRequest_Mutation, error) {
	return &pb.MutateFieldRequest_Mutation{
		Mutation: &pb.MutateFieldRequest_Mutation_Name{
			Name: string(n),
		},
	}, nil
}

// FieldTypeMutation creates a schema field mutation which changes the type of a field.
func FieldTypeMutation(ty FieldType) FieldMutation {
	return fieldTypeMutation(ty)
}

type fieldTypeMutation FieldType

func (t fieldTypeMutation) proto() (*pb.MutateFieldRequest_Mutation, error) {
	ty, err := FieldType(t).proto()
	if err != nil {
		return nil, err
	}
	return &pb.MutateFieldRequest_Mutation{
		Mutation: &pb.MutateFieldRequest_Mutation_Type{
			Type: ty,
		},
	}, nil
}

// FieldModeMutation creates a schema field mutation which changes the unique constraint on a field.
func FieldModeMutation(m FieldMode) FieldMutation {
	return fieldModeMutation(m)
}

type fieldModeMutation FieldMode

func (m fieldModeMutation) proto() (*pb.MutateFieldRequest_Mutation, error) {
	pbm, err := FieldMode(m).proto()
	if err != nil {
		return nil, err
	}

	return &pb.MutateFieldRequest_Mutation{
		Mutation: &pb.MutateFieldRequest_Mutation_Mode{
			Mode: pbm,
		},
	}, nil
}

// FieldAddIndexMutation adds a schema field mutation which adds an index to a field.
func FieldAddIndexMutation(x FieldIndex) FieldMutation {
	return fieldAddIndexMutation(x)
}

type fieldAddIndexMutation FieldIndex

func (i fieldAddIndexMutation) proto() (*pb.MutateFieldRequest_Mutation, error) {
	return &pb.MutateFieldRequest_Mutation{
		Mutation: &pb.MutateFieldRequest_Mutation_AddIndex{
			AddIndex: FieldIndex(i).proto(),
		},
	}, nil
}

// FieldRepeatedMutation creates a schema field mutation which changes the repeated property on a field.
func FieldRepeatedMutation(repeated bool) FieldMutation {
	return fieldRepeatedMutation(repeated)
}

type fieldRepeatedMutation bool

func (u fieldRepeatedMutation) proto() (*pb.MutateFieldRequest_Mutation, error) {
	return &pb.MutateFieldRequest_Mutation{
		Mutation: &pb.MutateFieldRequest_Mutation_Repeated{
			Repeated: bool(u),
		},
	}, nil
}

// FieldMutation is an interface which is satisfied by schema field mutations.
type FieldMutation interface {
	proto() (*pb.MutateFieldRequest_Mutation, error)
}

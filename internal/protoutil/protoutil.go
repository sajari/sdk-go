package protoutil

import (
	"fmt"
	"strconv"
	"time"

	enginev2pb "code.sajari.com/protogen-go/sajari/engine/v2"
	structpb "github.com/golang/protobuf/ptypes/struct"
)

func FromProto(v *enginev2pb.Value) (interface{}, error) {
	switch v := v.Value.(type) {
	case *enginev2pb.Value_Single:
		return v.Single, nil

	case *enginev2pb.Value_Repeated_:
		return v.Repeated.Values, nil

	default:
		return nil, fmt.Errorf("unexpected type: %T", v)
	}
}

func Single(x interface{}) (*enginev2pb.Value, error) {
	switch x := x.(type) {
	case int, uint, int64, uint64, int32, uint32, int16, uint16,
		int8, uint8, float32, float64, string, bool:
		return &enginev2pb.Value{
			Value: &enginev2pb.Value_Single{
				Single: fmt.Sprintf("%v", x),
			},
		}, nil

	default:
		return nil, fmt.Errorf("expected single value, got %T", x)
	}
}

func Value(x interface{}) (*enginev2pb.Value, error) {
	switch x := x.(type) {
	case string:
		return &enginev2pb.Value{
			Value: &enginev2pb.Value_Single{
				Single: x,
			},
		}, nil

	case int, uint, int64, uint64, int32, uint32, int16, uint16,
		int8, uint8, float32, float64, bool:
		return &enginev2pb.Value{
			Value: &enginev2pb.Value_Single{
				Single: fmt.Sprintf("%v", x),
			},
		}, nil

	case time.Time:
		return &enginev2pb.Value{
			Value: &enginev2pb.Value_Single{
				Single: strconv.FormatInt(x.Unix(), 10),
			},
		}, nil
	}

	var vs []string
	switch x := x.(type) {
	case []string:
		vs = x

	case []int:
		vs = make([]string, 0, len(x))
		for _, v := range x {
			vs = append(vs, strconv.FormatInt(int64(v), 10))
		}

	case []int64:
		vs = make([]string, 0, len(x))
		for _, v := range x {
			vs = append(vs, strconv.FormatInt(v, 10))
		}

	case []float32:
		vs = make([]string, 0, len(x))
		for _, v := range x {
			vs = append(vs, strconv.FormatFloat(float64(v), 'g', -1, 32))
		}

	case []float64:
		vs = make([]string, 0, len(x))
		for _, v := range x {
			vs = append(vs, strconv.FormatFloat(v, 'g', -1, 64))
		}

	case []interface{}:
		vs = make([]string, 0, len(x))
		for _, v := range x {
			vs = append(vs, fmt.Sprintf("%v", v))
		}

	default:
		return nil, fmt.Errorf("unsupported value: %T", x)
	}

	return &enginev2pb.Value{
		Value: &enginev2pb.Value_Repeated_{
			Repeated: &enginev2pb.Value_Repeated{
				Values: vs,
			},
		},
	}, nil
}

func Values(m map[string]interface{}) (map[string]*enginev2pb.Value, error) {
	values := make(map[string]*enginev2pb.Value, len(m))
	for k, v := range m {
		vv, err := Value(v)
		if err != nil {
			return nil, fmt.Errorf("could not convert value for %v: %w", k, err)
		}
		values[k] = vv
	}
	return values, nil
}

func Map(st *structpb.Struct) (map[string]string, error) {
	out := make(map[string]string, len(st.GetFields()))
	for k, v := range st.GetFields() {
		_, ok := v.Kind.(*structpb.Value_StringValue)
		if !ok {
			return nil, fmt.Errorf("non-string value (%T)", v)
		}
		out[k] = v.GetStringValue()
	}
	return out, nil
}

func Struct(m map[string]string) *structpb.Struct {
	fields := make(map[string]*structpb.Value, len(m))
	for k, v := range m {
		fields[k] = &structpb.Value{
			Kind: &structpb.Value_StringValue{
				StringValue: v,
			},
		}
	}

	return &structpb.Struct{
		Fields: fields,
	}
}

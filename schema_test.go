package sajari

import (
	"context"
	"errors"
	"testing"
)

func TestFieldIterator_Next_ErrDone(t *testing.T) {
	it := &FieldIterator{
		ctx:    context.Background(),
		fields: []Field{},
		end:    true,
	}
	_, err := it.Next()
	if err == ErrDone {
		t.Errorf("Next() error = ErrDone, want wrapped ErrDone")
	}
	if !errors.Is(err, ErrDone) {
		t.Errorf("Next() error = %v, want wrapped ErrDone", err)
	}
}

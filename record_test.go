package sajari

import (
	"context"
	"errors"
	"testing"
)

func TestKeyIterator_Next_Done(t *testing.T) {
	it := &KeyIterator{
		ctx:  context.Background(),
		keys: []*Key{},
		end:  true,
	}
	_, err := it.Next()
	if err == ErrDone {
		t.Errorf("Next() error = ErrDone, want wrapped ErrDone")
	}
	if !errors.Is(err, ErrDone) {
		t.Errorf("Next() error = %v, want wrapped ErrDone", err)
	}
}

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

func TestKey_String(t *testing.T) {
	tests := []struct {
		key  *Key
		want string
	}{
		{NewKey("id", 1234), "Key{Field: id, Value: 1234}"},
		{NewKey("id", "1234"), "Key{Field: id, Value: 1234}"},
		{NewKey("url", "https://example.com"), "Key{Field: url, Value: https://example.com}"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.key.String(); got != tt.want {
				t.Errorf("Key.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

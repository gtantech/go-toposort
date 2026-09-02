package vertex

import (
	"testing"
)

func TestValue(t *testing.T) {
	v := New(1)

	if want := 1; v.Value() != want {
		t.Errorf("want %v, got %v", want, v.Value())
	}
}

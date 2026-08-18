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

func TestIsVisitedDefault(t *testing.T) {
	v := New(1)

	if want := false; v.IsVisited() != want {
		t.Errorf("want %v, got %v", want, v.IsVisited())
	}
}

func TestSetVisited(t *testing.T) {
	v := New(1)
	v.SetVisited()
	if want := true; v.IsVisited() != want {
		t.Errorf("want %v, got %v", want, v.IsVisited())
	}
}

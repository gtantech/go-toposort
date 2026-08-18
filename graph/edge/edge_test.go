package edge

import (
	"testing"

	"github.com/gtantech/go-toposort/graph/vertex"
)

func TestValue(t *testing.T) {
	v1 := vertex.New("v1")
	v2 := vertex.New("v2")
	e := New(1, v1, v2)

	if want := 1; want != e.Value() {
		t.Errorf("want %v, got %v", want, e.Value())
	}
}

func TestGetOrigin(t *testing.T) {
	v1 := vertex.New("v1")
	v2 := vertex.New("v2")
	e := New(1, v1, v2)

	if want := v1; want != e.GetOrigin() {
		t.Errorf("want %v, got %v", want, e.GetOrigin())
	}
}

func TestGetDestination(t *testing.T) {
	v1 := vertex.New("v1")
	v2 := vertex.New("v2")
	e := New(1, v1, v2)

	if want := v2; want != e.GetDestination() {
		t.Errorf("want %v, got %v", want, e.GetDestination())
	}
}

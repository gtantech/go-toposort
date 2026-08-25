package toposort

import (
	"fmt"
	"testing"

	"github.com/gtantech/toposort/graph/vertex"
)

func TestErrorMsg(t *testing.T) {
	v1 := vertex.New("1")
	v2 := vertex.New("2")

	err := CycleDetectedError[string, string]{"1-2", v1, v2}

	want := fmt.Sprintf("encountered cycle in graph for edge: %v from %v to %v", "1-2", "1", "2")
	if got := err.Error(); got != want {
		t.Errorf("unexpected error message. got: %v, want %v", got, want)
	}
}

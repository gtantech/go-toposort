package graph

import (
	"slices"
	"testing"

	"github.com/gtantech/go-toposort/graph/edge"
	"github.com/gtantech/go-toposort/graph/vertex"
)

func TestNew(t *testing.T) {
	g := New[int, string]()
	if got := len(g.Vertices()); got != 0 {
		t.Errorf("expected zero size verticies, got %v", got)
	}
}

func TestAddEdges(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)
	two := vertex.New(2)
	e1 := edge.New("1-2", one, two)
	g.AddEdge(e1)
	if got := len(g.OutgoingEdges(one)); got != 1 {
		t.Errorf("expected number of edges to be 1, got %v", got)
	}

	if got := g.OutgoingEdges(one); got[0] != e1 {
		t.Errorf("unexpected edge, got %v", e1.Value())
	}

	if got := len(g.Vertices()); got != 2 {
		t.Errorf("expected number of vertices to be 2, got %v", got)
	}

	if !slices.Contains(g.Vertices(), one) {
		t.Errorf("missing vertex %v in Verticies()", one)
	}

	if !slices.Contains(g.Vertices(), two) {
		t.Errorf("missing vertex %v in Verticies()", two)
	}
}

func TestAddSameOriginEdges(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)
	two := vertex.New(2)
	three := vertex.New(3)
	e1 := edge.New("1-2", one, two)
	e2 := edge.New("1-3", one, three)
	g.AddEdge(e1)
	g.AddEdge(e2)
	if got := len(g.OutgoingEdges(one)); got != 2 {
		t.Errorf("expected number of edges to be 2, got %v", got)
	}

	if got := g.OutgoingEdges(one); got[0] != e1 {
		t.Errorf("unexpected edge, want %v got %v", e1, got[0])
	}

	if got := g.OutgoingEdges(one); got[1] != e2 {
		t.Errorf("unexpected edge, want %v got %v", e2, got[1])
	}

	if got := len(g.Vertices()); got != 3 {
		t.Errorf("expected number of vertices to be 3, got %v", got)
	}

	if !slices.Contains(g.Vertices(), one) {
		t.Errorf("missing vertex %v in Verticies()", one)
	}

	if !slices.Contains(g.Vertices(), two) {
		t.Errorf("missing vertex %v in Verticies()", two)
	}

	if !slices.Contains(g.Vertices(), three) {
		t.Errorf("missing vertex %v in Verticies()", three)
	}
}

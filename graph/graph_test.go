package graph

import (
	"reflect"
	"slices"
	"testing"

	"github.com/gtantech/toposort/graph/edge"
	"github.com/gtantech/toposort/graph/vertex"
)

func TestNew(t *testing.T) {
	g := New[int, string]()
	if got := len(g.Vertices()); got != 0 {
		t.Errorf("expected zero size verticies, got %v", got)
	}
}

func TestAddVertex(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)

	if len(g.Vertices()) != 0 {
		t.Error("expected empty graph")
	}

	g.AddVertex(one)

	if !slices.Contains(g.Vertices(), one) {
		t.Errorf("missing vertex %v in Verticies()", one)
	}

}

func TestAddEdges(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)
	two := vertex.New(2)
	e1 := edge.New("1-2", one, two)
	g.AddEdge(e1)
	if got := len(slices.Collect(g.OutgoingVertices(one))); got != 1 {
		t.Errorf("expected number of edges to be 1, got %v", got)
	}

	if got := slices.Collect(g.OutgoingVertices(one)); got[0] != two {
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
	if got := len(slices.Collect(g.OutgoingVertices(one))); got != 2 {
		t.Errorf("expected number of edges to be 2, got %v", got)
	}

	if got := slices.Collect(g.OutgoingVertices(one)); !(reflect.DeepEqual(got, []vertex.Vertex[int]{two, three}) || reflect.DeepEqual(got, []vertex.Vertex[int]{three, two})) {
		t.Errorf("unexpected edge, want %v got %v", []edge.Edge[string, int]{e1, e2}, got)
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

func TestRemoveEdge(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)
	two := vertex.New(2)
	e1 := edge.New("1-2", one, two)
	g.AddEdge(e1)
	if got := len(slices.Collect(g.OutgoingVertices(one))); got != 1 {
		t.Errorf("expected number of edges to be 1, got %v", got)
	}

	if got := slices.Collect(g.OutgoingVertices(one)); got[0] != two {
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

	g.RemoveEdge(e1)

	if got := len(g.Vertices()); got != 2 {
		t.Errorf("expected number of vertices to be 2, got %v", got)
	}

	if got := len(slices.Collect(g.OutgoingVertices(one))); got != 0 {
		t.Errorf("expected number of edges to be 1, got %v", got)
	}
}

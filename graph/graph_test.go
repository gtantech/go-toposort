package graph

import (
	"reflect"
	"slices"
	"testing"

	"github.com/gtantech/toposort/graph/vertex"
)

func TestNew(t *testing.T) {
	g := New[int, string]()
	if got := len(slices.Collect(g.Vertices())); got != 0 {
		t.Errorf("expected zero size verticies, got %v", got)
	}
}

func TestAddVertex(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)

	if len(slices.Collect(g.Vertices())) != 0 {
		t.Error("expected empty graph")
	}

	g.AddVertex(one)

	if !slices.Contains(slices.Collect(g.Vertices()), one) {
		t.Errorf("missing vertex %v in Verticies()", one)
	}

}

func TestAddEdges(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)
	two := vertex.New(2)
	value := "1-2"
	g.AddEdge(value, one, two)
	if got := len(slices.Collect(g.OutgoingVertices(one))); got != 1 {
		t.Errorf("expected number of edges to be 1, got %v", got)
	}

	if got := slices.Collect(g.OutgoingVertices(one)); got[0] != two {
		t.Errorf("unexpected edge, got %v", value)
	}

	if got := len(slices.Collect(g.Vertices())); got != 2 {
		t.Errorf("expected number of vertices to be 2, got %v", got)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), one) {
		t.Errorf("missing vertex %v in Verticies()", one)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), two) {
		t.Errorf("missing vertex %v in Verticies()", two)
	}
}

func TestAddSameOriginEdges(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)
	two := vertex.New(2)
	three := vertex.New(3)
	e1 := "1-2"
	e2 := "1-3"
	g.AddEdge(e1, one, two)
	g.AddEdge(e2, one, three)
	if got := len(slices.Collect(g.OutgoingVertices(one))); got != 2 {
		t.Errorf("expected number of edges to be 2, got %v", got)
	}

	if got := slices.Collect(g.OutgoingVertices(one)); !(reflect.DeepEqual(got, []vertex.Vertex[int]{two, three}) || reflect.DeepEqual(got, []vertex.Vertex[int]{three, two})) {
		t.Errorf("unexpected edge, want %v got %v", []vertex.Vertex[int]{two, three}, got)
	}

	if got := len(slices.Collect(g.Vertices())); got != 3 {
		t.Errorf("expected number of vertices to be 3, got %v", got)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), one) {
		t.Errorf("missing vertex %v in Verticies()", one)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), two) {
		t.Errorf("missing vertex %v in Verticies()", two)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), three) {
		t.Errorf("missing vertex %v in Verticies()", three)
	}
}

func TestRemoveEdge(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)
	two := vertex.New(2)
	e1 := "1-2"
	g.AddEdge(e1, one, two)
	if got := len(slices.Collect(g.OutgoingVertices(one))); got != 1 {
		t.Errorf("expected number of edges to be 1, got %v", got)
	}

	if got := slices.Collect(g.OutgoingVertices(one)); got[0] != two {
		t.Errorf("unexpected edge, got %v", e1)
	}

	if got := len(slices.Collect(g.Vertices())); got != 2 {
		t.Errorf("expected number of vertices to be 2, got %v", got)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), one) {
		t.Errorf("missing vertex %v in Verticies()", one)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), two) {
		t.Errorf("missing vertex %v in Verticies()", two)
	}

	g.RemoveEdge(one, two)

	if got := len(slices.Collect(g.Vertices())); got != 2 {
		t.Errorf("expected number of vertices to be 2, got %v", got)
	}

	if got := len(slices.Collect(g.OutgoingVertices(one))); got != 0 {
		t.Errorf("expected number of edges to be 1, got %v", got)
	}
}

func TestGetEdgeValue(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)
	two := vertex.New(2)
	e1 := "1-2"
	g.AddEdge(e1, one, two)

	ev, ok := g.GetEdgeValue(one, two)

	if !ok {
		t.Errorf("expected edge value to be ok, got not ok")
	}

	if ev != e1 {
		t.Errorf("unexpected value, got %v want %v", ev, e1)
	}
}

func TestGetEdgeValueNoValueFound(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)
	two := vertex.New(2)
	three := vertex.New(3)
	e1 := "1-2"
	g.AddEdge(e1, one, two)

	ev, ok := g.GetEdgeValue(one, three)

	if ok {
		t.Errorf("expected edge value to be not ok, got ok")
	}

	if ev != "" {
		t.Errorf("unexpected value, got %v want %v", ev, "")
	}

	ev, ok = g.GetEdgeValue(two, three)

	if ok {
		t.Errorf("expected edge value to be not ok, got ok")
	}

	if ev != "" {
		t.Errorf("unexpected value, got %v want %v", ev, "")
	}
}

func TestRemoveVertex(t *testing.T) {
	g := New[int, string]()
	one := vertex.New(1)
	two := vertex.New(2)
	e1 := "1-2"
	g.AddEdge(e1, one, two)

	g.RemoveVertex(one)
	if _, ok := g.GetEdgeValue(one, two); ok {
		t.Errorf("expected edge to be removed")
	}

	for v := range g.Vertices() {
		if v == one {
			t.Errorf("expected vertex %v to be removed", v)
		}
	}

	g.AddEdge(e1, one, two)

	g.RemoveVertex(two)
	if _, ok := g.GetEdgeValue(one, two); ok {
		t.Errorf("expected edge to be removed")
	}

	for v := range g.Vertices() {
		if v == two {
			t.Errorf("expected vertex %v to be removed", v)
		}
	}
}

package graph

import (
	"reflect"
	"slices"
	"testing"
)

func TestNew(t *testing.T) {
	g := New[int, string]()
	if got := len(slices.Collect(g.Vertices())); got != 0 {
		t.Errorf("expected zero size verticies, got %v", got)
	}
}

func TestAddVertex(t *testing.T) {
	g := New[int, string]()

	if len(slices.Collect(g.Vertices())) != 0 {
		t.Error("expected empty graph")
	}

	g.AddVertex(1)

	if !slices.Contains(slices.Collect(g.Vertices()), 1) {
		t.Errorf("missing vertex %v in Verticies()", 1)
	}

}

func TestAddEdges(t *testing.T) {
	g := New[int, string]()

	value := "1-2"
	g.AddEdge(value, 1, 2)
	if got := len(slices.Collect(g.OutgoingVertices(1))); got != 1 {
		t.Errorf("expected number of edges to be 1, got %v", got)
	}

	if got := slices.Collect(g.OutgoingVertices(1)); got[0] != 2 {
		t.Errorf("unexpected edge, got %v", value)
	}

	if got := len(slices.Collect(g.Vertices())); got != 2 {
		t.Errorf("expected number of vertices to be 2, got %v", got)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), 1) {
		t.Errorf("missing vertex %v in Verticies()", 1)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), 2) {
		t.Errorf("missing vertex %v in Verticies()", 2)
	}
}

func TestIncomingVertices(t *testing.T) {
	g := New[int, string]()

	g.AddEdge("1-2", 1, 2)
	g.AddEdge("3-2", 3, 2)
	if got := len(slices.Collect(g.IncomingVertices(2))); got != 2 {
		t.Errorf("expected number of edges to be 2, got %v", got)
	}

	if got := slices.Collect(g.IncomingVertices(2)); !(slices.Contains(got, 1) && slices.Contains(got, 3) && !slices.Contains(got, 2) && len(got) == 2) {
		t.Errorf("unexpected edge, got %v", got)
	}
}

func TestAddSameOriginEdges(t *testing.T) {
	g := New[int, string]()

	e1 := "1-2"
	e2 := "1-3"
	g.AddEdge(e1, 1, 2)
	g.AddEdge(e2, 1, 3)
	if got := len(slices.Collect(g.OutgoingVertices(1))); got != 2 {
		t.Errorf("expected number of edges to be 2, got %v", got)
	}

	if got := slices.Collect(g.OutgoingVertices(1)); !(reflect.DeepEqual(got, []int{2, 3}) || reflect.DeepEqual(got, []int{3, 2})) {
		t.Errorf("unexpected edge, want %v got %v", []int{2, 3}, got)
	}

	if got := len(slices.Collect(g.Vertices())); got != 3 {
		t.Errorf("expected number of vertices to be 3, got %v", got)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), 1) {
		t.Errorf("missing vertex %v in Verticies()", 1)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), 2) {
		t.Errorf("missing vertex %v in Verticies()", 2)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), 3) {
		t.Errorf("missing vertex %v in Verticies()", 3)
	}
}

func TestRemoveEdge(t *testing.T) {
	g := New[int, string]()

	e1 := "1-2"
	g.AddEdge(e1, 1, 2)
	if got := len(slices.Collect(g.OutgoingVertices(1))); got != 1 {
		t.Errorf("expected number of edges to be 1, got %v", got)
	}

	if got := slices.Collect(g.OutgoingVertices(1)); got[0] != 2 {
		t.Errorf("unexpected edge, got %v", e1)
	}

	if got := len(slices.Collect(g.Vertices())); got != 2 {
		t.Errorf("expected number of vertices to be 2, got %v", got)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), 1) {
		t.Errorf("missing vertex %v in Verticies()", 1)
	}

	if !slices.Contains(slices.Collect(g.Vertices()), 2) {
		t.Errorf("missing vertex %v in Verticies()", 2)
	}

	g.RemoveEdge(1, 2)

	if got := len(slices.Collect(g.Vertices())); got != 2 {
		t.Errorf("expected number of vertices to be 2, got %v", got)
	}

	if got := len(slices.Collect(g.OutgoingVertices(1))); got != 0 {
		t.Errorf("expected number of edges to be 1, got %v", got)
	}
}

func TestGetEdgeValue(t *testing.T) {
	g := New[int, string]()

	e1 := "1-2"
	g.AddEdge(e1, 1, 2)

	ev, ok := g.GetEdgeValue(1, 2)

	if !ok {
		t.Errorf("expected edge value to be ok, got not ok")
	}

	if ev != e1 {
		t.Errorf("unexpected value, got %v want %v", ev, e1)
	}
}

func TestGetEdgeValueNoValueFound(t *testing.T) {
	g := New[int, string]()

	e1 := "1-2"
	g.AddEdge(e1, 1, 2)

	ev, ok := g.GetEdgeValue(1, 3)

	if ok {
		t.Errorf("expected edge value to be not ok, got ok")
	}

	if ev != "" {
		t.Errorf("unexpected value, got %v want %v", ev, "")
	}

	ev, ok = g.GetEdgeValue(2, 3)

	if ok {
		t.Errorf("expected edge value to be not ok, got ok")
	}

	if ev != "" {
		t.Errorf("unexpected value, got %v want %v", ev, "")
	}
}

func TestRemoveVertex(t *testing.T) {
	g := New[int, string]()

	e1 := "1-2"
	g.AddEdge(e1, 1, 2)

	g.RemoveVertex(1)
	if _, ok := g.GetEdgeValue(1, 2); ok {
		t.Errorf("expected edge to be removed")
	}

	for v := range g.Vertices() {
		if v == 1 {
			t.Errorf("expected vertex %v to be removed", v)
		}
	}

	g.AddEdge(e1, 1, 2)

	g.RemoveVertex(2)
	if _, ok := g.GetEdgeValue(1, 2); ok {
		t.Errorf("expected edge to be removed")
	}

	for v := range g.Vertices() {
		if v == 2 {
			t.Errorf("expected vertex %v to be removed", v)
		}
	}
}

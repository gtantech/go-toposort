package toposort

import (
	"fmt"
	"testing"

	"github.com/gtantech/toposort/graph"
	"github.com/gtantech/toposort/stack"
)

type mockDfsTopoStack[T comparable] struct {
	stack.Stack[T]
}

func (s *mockDfsTopoStack[T]) Push(v T) error {
	s.Stack.Push(v)
	return fmt.Errorf("sent error to panic")
}

type mockGraph[V comparable, E any] struct {
	graph.Graph[V, E]
}

func (g *mockGraph[V, E]) GetEdgeValue(origin V, destination V) (E, bool) {
	var zero E
	return zero, false
}

func TestDfsTopoUnhandledStackDuplicateValuesError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("unhandled error did not panic")
		}
	}()

	DAG := graph.New[string, string]()
	s := mockDfsTopoStack[string]{Stack: stack.New[string]()}
	DAG.AddEdge("AB", "A", "B")
	isVisited := make(map[string]bool)
	dfsTopo(DAG, "A", &s, isVisited)
}

func TestDfsFailedToGetEdgeValue(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("failed to get edge did not panic")
		}
	}()
	DAG := mockGraph[string, string]{Graph: graph.New[string, string]()}
	s := stack.New[string]()

	DAG.AddEdge("AB", "A", "B")
	DAG.AddEdge("BA", "B", "A")
	isVisited := make(map[string]bool)
	dfsTopo(&DAG, "A", s, isVisited)
}

func TestSort(t *testing.T) {

	DAG := graph.New[string, string]()

	DAG.AddEdge("AB", "A", "B")
	DAG.AddEdge("AC", "A", "C")
	DAG.AddEdge("BC", "B", "C")
	DAG.AddEdge("BD", "B", "D")
	DAG.AddEdge("CE", "C", "E")
	DAG.AddEdge("ED", "E", "D")
	DAG.AddEdge("EF", "E", "F")

	order, err := TopologicalSort(DAG)
	if err != nil {
		t.Errorf("unexpected error occurred. Error: %v", err)
	}
	isVisited := make(map[string]bool)
	for v := range DAG.Vertices() {
		isVisited[v] = false
	}

	for _, v := range order {
		isVisited[v] = true
		for destination := range DAG.OutgoingVertices(v) {
			if isVisited[destination] {
				t.Fatalf("%v was visited before the current vertex %v", destination, v)
			}
		}
	}
}

func TestSortWithCycle(t *testing.T) {

	DAG := graph.New[string, string]()

	DAG.AddEdge("AB", "A", "B")
	DAG.AddEdge("AC", "A", "C")
	DAG.AddEdge("BC", "B", "C")
	DAG.AddEdge("BD", "B", "D")
	DAG.AddEdge("CE", "C", "E")
	DAG.AddEdge("ED", "E", "D")
	DAG.AddEdge("EF", "E", "F")
	DAG.AddEdge("FA", "F", "A") //add a return edge from F to A to add a cycle

	_, err := TopologicalSort(DAG)
	if err == nil {
		t.Errorf("expected error")
	}
}

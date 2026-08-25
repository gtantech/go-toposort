package toposort

import (
	"fmt"
	"testing"

	"github.com/gtantech/toposort/graph"
	"github.com/gtantech/toposort/graph/edge"
	"github.com/gtantech/toposort/graph/vertex"
	"github.com/gtantech/toposort/stack"
)

type mockVertex[V any] struct {
	value     V
	isVisited bool
}

func (v *mockVertex[V]) String() string {
	return fmt.Sprintf("%v", v.value)
}

// IsVisited implements [graph.Vertex].
func (v *mockVertex[V]) IsVisited() bool {
	return v.isVisited
}

// SetVisited implements [graph.Vertex].
func (v *mockVertex[V]) SetVisited() {
	v.isVisited = true
}

// Value implements [graph.Vertex].
func (v *mockVertex[V]) Value() V {
	return v.value
}

type mockDfsTopoStack[T comparable] struct {
	stack.Stack[vertex.Vertex[T]]
}

func (s *mockDfsTopoStack[T]) Push(v vertex.Vertex[T]) error {
	s.Stack.Push(v)
	return fmt.Errorf("sent error to panic")
}

func TestDfsTopoUnhandledStackDuplicateValuesError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("unhandled error did not panic")
		}
	}()

	DAG := graph.New[string, string]()
	s := mockDfsTopoStack[string]{Stack: stack.New[vertex.Vertex[string]]()}
	A := mockVertex[string]{value: "A"}
	B := mockVertex[string]{value: "B"}
	DAG.AddEdge(edge.New("AB", &A, &B))
	dfsTopo(DAG, &A, &s)
}

func TestSort(t *testing.T) {
	A := mockVertex[string]{value: "A"}
	B := mockVertex[string]{value: "B"}
	C := mockVertex[string]{value: "C"}
	D := mockVertex[string]{value: "D"}
	E := mockVertex[string]{value: "E"}
	F := mockVertex[string]{value: "F"}

	DAG := graph.New[string, string]()

	DAG.AddEdge(edge.New("AB", &A, &B))
	DAG.AddEdge(edge.New("AC", &A, &C))
	DAG.AddEdge(edge.New("BC", &B, &C))
	DAG.AddEdge(edge.New("BD", &B, &D))
	DAG.AddEdge(edge.New("CE", &C, &E))
	DAG.AddEdge(edge.New("ED", &E, &D))
	DAG.AddEdge(edge.New("EF", &E, &F))

	order, err := TopologicalSort(DAG)
	if err != nil {
		t.Errorf("unexpected error occurred. Error: %v", err)
	}
	A.isVisited = false
	B.isVisited = false
	C.isVisited = false
	D.isVisited = false
	E.isVisited = false
	F.isVisited = false

	for _, v := range order {
		v.SetVisited()
		for destination := range DAG.OutgoingVertices(v) {
			if destination.IsVisited() {
				t.Fatalf("%v was visited before the current vertex %v", destination, v)
			}
		}
	}
}

func TestSortWithCycle(t *testing.T) {
	A := mockVertex[string]{value: "A"}
	B := mockVertex[string]{value: "B"}
	C := mockVertex[string]{value: "C"}
	D := mockVertex[string]{value: "D"}
	E := mockVertex[string]{value: "E"}
	F := mockVertex[string]{value: "F"}

	DAG := graph.New[string, string]()

	DAG.AddEdge(edge.New("AB", &A, &B))
	DAG.AddEdge(edge.New("AC", &A, &C))
	DAG.AddEdge(edge.New("BC", &B, &C))
	DAG.AddEdge(edge.New("BD", &B, &D))
	DAG.AddEdge(edge.New("CE", &C, &E))
	DAG.AddEdge(edge.New("ED", &E, &D))
	DAG.AddEdge(edge.New("EF", &E, &F))
	DAG.AddEdge(edge.New("FA", &F, &A)) //add a return edge from F to A to add a cycle

	_, err := TopologicalSort(DAG)
	if err == nil {
		t.Errorf("expected error")
	}
}

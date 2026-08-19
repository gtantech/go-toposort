package toposort

import (
	"fmt"
	"testing"

	"github.com/gtantech/toposort/graph"
	"github.com/gtantech/toposort/graph/edge"
)

type testVertex[V any] struct {
	value     V
	isVisited bool
}

func (v *testVertex[V]) String() string {
	return fmt.Sprintf("%v", v.value)
}

// IsVisited implements [graph.Vertex].
func (v *testVertex[V]) IsVisited() bool {
	return v.isVisited
}

// SetVisited implements [graph.Vertex].
func (v *testVertex[V]) SetVisited() {
	v.isVisited = true
}

// Value implements [graph.Vertex].
func (v *testVertex[V]) Value() V {
	return v.value
}

func TestSort(t *testing.T) {
	A := testVertex[string]{value: "A"}
	B := testVertex[string]{value: "B"}
	C := testVertex[string]{value: "C"}
	D := testVertex[string]{value: "D"}
	E := testVertex[string]{value: "E"}
	F := testVertex[string]{value: "F"}

	DAG := graph.New[string, string]()

	DAG.AddEdge(edge.New("AB", &A, &B))
	DAG.AddEdge(edge.New("AC", &A, &C))
	DAG.AddEdge(edge.New("BC", &B, &C))
	DAG.AddEdge(edge.New("BD", &B, &D))
	DAG.AddEdge(edge.New("CE", &C, &E))
	DAG.AddEdge(edge.New("ED", &E, &D))
	DAG.AddEdge(edge.New("EF", &E, &F))

	order := TopologicalSort(DAG)
	A.isVisited = false
	B.isVisited = false
	C.isVisited = false
	D.isVisited = false
	E.isVisited = false
	F.isVisited = false

	for _, v := range order {
		v.SetVisited()
		for _, e := range DAG.OutgoingEdges(v) {
			if e.GetDestination().IsVisited() {
				t.Fatalf("%v was visited before the current vertex %v", e.GetDestination(), v)
			}
		}
	}
}

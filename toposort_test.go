package toposort

import (
	"fmt"
	"testing"

	"github.com/gtantech/go-toposort/graph"
)

var _ graph.Vertex[int] = (*vertex[int])(nil) //ensures queue implements Queue at compile time
type vertex[V any] struct {
	value     V
	isVisited bool
}

func (v *vertex[V]) String() string {
	return fmt.Sprintf("%v", v.value)
}

// IsVisited implements [graph.Vertex].
func (v *vertex[V]) IsVisited() bool {
	return v.isVisited
}

// SetVisited implements [graph.Vertex].
func (v *vertex[V]) SetVisited() {
	v.isVisited = true
}

// Value implements [graph.Vertex].
func (v *vertex[V]) Value() V {
	return v.value
}

var _ graph.Edge[string, string] = (*edge[string, string])(nil)

type edge[E any, V any] struct {
	destination graph.Vertex[V]
	name        string
}

func (e *edge[E, V]) String() string {
	return e.name
}

// GetDestination implements [graph.Edge].
func (e *edge[E, V]) GetDestination() graph.Vertex[V] {
	return e.destination
}

var _ graph.Graph[string, string] = (*dag[string, string])(nil)

type dag[V comparable, E any] struct {
	vertices      []graph.Vertex[V]
	outgoingEdges map[graph.Vertex[V]][]graph.Edge[E, V]
}

func (d *dag[V, E]) AddVertex(v graph.Vertex[V]) {
	d.vertices = append(d.vertices, v)
}

func (d *dag[V, E]) AddEdge(origin graph.Vertex[V], e graph.Edge[E, V]) {
	value, ok := d.outgoingEdges[origin]
	if !ok {
		d.outgoingEdges[origin] = []graph.Edge[E, V]{e}
	} else {
		d.outgoingEdges[origin] = append(value, e)
	}
}

// OutgoingEdges implements [graph.Graph].
func (d *dag[V, E]) OutgoingEdges(vertex graph.Vertex[V]) []graph.Edge[E, V] {
	return d.outgoingEdges[vertex]
}

// Vertices implements [graph.Graph].
func (d *dag[V, E]) Vertices() []graph.Vertex[V] {
	return d.vertices
}

func TestSort(t *testing.T) {
	A := vertex[string]{value: "A"}
	B := vertex[string]{value: "B"}
	C := vertex[string]{value: "C"}
	D := vertex[string]{value: "D"}
	E := vertex[string]{value: "E"}
	F := vertex[string]{value: "F"}

	DAG := dag[string, int]{vertices: []graph.Vertex[string]{}, outgoingEdges: make(map[graph.Vertex[string]][]graph.Edge[int, string])}

	DAG.AddVertex(&A)
	DAG.AddVertex(&B)
	DAG.AddVertex(&C)
	DAG.AddVertex(&D)
	DAG.AddVertex(&E)
	DAG.AddVertex(&F)

	DAG.AddEdge(&A, &edge[int, string]{destination: &B, name: "AB"})
	DAG.AddEdge(&A, &edge[int, string]{destination: &C, name: "AC"})
	DAG.AddEdge(&B, &edge[int, string]{destination: &C, name: "BC"})
	DAG.AddEdge(&B, &edge[int, string]{destination: &D, name: "BD"})
	DAG.AddEdge(&C, &edge[int, string]{destination: &E, name: "CE"})
	DAG.AddEdge(&E, &edge[int, string]{destination: &D, name: "ED"})
	DAG.AddEdge(&E, &edge[int, string]{destination: &F, name: "EF"})

	order := TopologicalSort(&DAG)
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

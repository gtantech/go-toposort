package toposort

import (
	"fmt"
	"testing"

	"github.com/gtantech/go-toposort/graph"
	"github.com/gtantech/go-toposort/graph/edge"
	"github.com/gtantech/go-toposort/graph/vertex"
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

var _ graph.Graph[string, string] = (*dag[string, string])(nil)

type dag[V comparable, E any] struct {
	vertices      []vertex.Vertex[V]
	outgoingEdges map[vertex.Vertex[V]][]edge.Edge[E, V]
}

func (d *dag[V, E]) AddVertex(v vertex.Vertex[V]) {
	d.vertices = append(d.vertices, v)
}

func (d *dag[V, E]) AddEdge(origin vertex.Vertex[V], e edge.Edge[E, V]) {
	value, ok := d.outgoingEdges[origin]
	if !ok {
		d.outgoingEdges[origin] = []edge.Edge[E, V]{e}
	} else {
		d.outgoingEdges[origin] = append(value, e)
	}
}

// OutgoingEdges implements [graph.Graph].
func (d *dag[V, E]) OutgoingEdges(vertex vertex.Vertex[V]) []edge.Edge[E, V] {
	return d.outgoingEdges[vertex]
}

// Vertices implements [graph.Graph].
func (d *dag[V, E]) Vertices() []vertex.Vertex[V] {
	return d.vertices
}

func TestSort(t *testing.T) {
	A := testVertex[string]{value: "A"}
	B := testVertex[string]{value: "B"}
	C := testVertex[string]{value: "C"}
	D := testVertex[string]{value: "D"}
	E := testVertex[string]{value: "E"}
	F := testVertex[string]{value: "F"}

	DAG := dag[string, int]{vertices: []vertex.Vertex[string]{}, outgoingEdges: make(map[vertex.Vertex[string]][]edge.Edge[int, string])}

	DAG.AddVertex(&A)
	DAG.AddVertex(&B)
	DAG.AddVertex(&C)
	DAG.AddVertex(&D)
	DAG.AddVertex(&E)
	DAG.AddVertex(&F)

	DAG.AddEdge(&A, edge.New("AB", &B))
	DAG.AddEdge(&A, edge.New("AC", &C))
	DAG.AddEdge(&B, edge.New("BC", &C))
	DAG.AddEdge(&B, edge.New("BD", &D))
	DAG.AddEdge(&C, edge.New("CE", &E))
	DAG.AddEdge(&E, edge.New("ED", &D))
	DAG.AddEdge(&E, edge.New("EF", &F))

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

package graph

import (
	"github.com/gtantech/go-toposort/graph/edge"
	"github.com/gtantech/go-toposort/graph/vertex"
)

type Graph[V any, E any] interface {
	OutgoingEdges(vertex vertex.Vertex[V]) []edge.Edge[E, V]
	Vertices() []vertex.Vertex[V]
	AddEdge(e edge.Edge[E, V])
}

var _ Graph[string, string] = (*dag[string, string])(nil)

type dag[V any, E any] struct {
	vertices        []vertex.Vertex[V]
	outgoingEdges   map[vertex.Vertex[V]][]edge.Edge[E, V]
	uniqueVerticies map[vertex.Vertex[V]]struct{}
}

func New[V any, E any]() Graph[V, E] {
	return &dag[V, E]{vertices: []vertex.Vertex[V]{}, outgoingEdges: make(map[vertex.Vertex[V]][]edge.Edge[E, V])}
}

func (d *dag[V, E]) AddEdge(e edge.Edge[E, V]) {
	value, ok := d.outgoingEdges[e.GetOrigin()]
	if !ok {
		d.outgoingEdges[e.GetOrigin()] = []edge.Edge[E, V]{e}
	} else {
		d.outgoingEdges[e.GetOrigin()] = append(value, e)
	}
	if _, ok := d.uniqueVerticies[e.GetOrigin()]; !ok {
		d.vertices = append(d.vertices, e.GetOrigin())
	}
	if _, ok := d.uniqueVerticies[e.GetDestination()]; !ok {
		d.vertices = append(d.vertices, e.GetDestination())
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

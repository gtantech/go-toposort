package graph

import (
	"maps"

	"github.com/gtantech/toposort/graph/edge"
	"github.com/gtantech/toposort/graph/vertex"
)

type Graph[V any, E any] interface {
	OutgoingVertices(vertex vertex.Vertex[V]) func(yield func(vertex.Vertex[V]) bool)
	Vertices() []vertex.Vertex[V]
	AddVertex(v vertex.Vertex[V])
	AddEdge(e edge.Edge[E, V])
	RemoveEdge(e edge.Edge[E, V])
}

var _ Graph[string, string] = (*dag[string, string])(nil)

type dag[V any, E any] struct {
	vertices           []vertex.Vertex[V]
	incomingToOutgoing map[vertex.Vertex[V]]map[vertex.Vertex[V]]E
	uniqueVerticies    map[vertex.Vertex[V]]struct{}
}

func New[V any, E any]() Graph[V, E] {
	return &dag[V, E]{vertices: []vertex.Vertex[V]{}, incomingToOutgoing: make(map[vertex.Vertex[V]]map[vertex.Vertex[V]]E), uniqueVerticies: make(map[vertex.Vertex[V]]struct{})}
}

func (d *dag[V, E]) AddVertex(v vertex.Vertex[V]) {
	if _, ok := d.uniqueVerticies[v]; !ok {
		d.uniqueVerticies[v] = struct{}{}
		d.vertices = append(d.vertices, v)
	}
}

func (d *dag[V, E]) RemoveEdge(e edge.Edge[E, V]) {
	outgoing, ok := d.incomingToOutgoing[e.GetOrigin()]
	if ok {
		//origin vertex exists
		delete(outgoing, e.GetDestination())
	}
}

func (d *dag[V, E]) AddEdge(e edge.Edge[E, V]) {
	outgoing, ok := d.incomingToOutgoing[e.GetOrigin()]
	if !ok {
		outgoing := map[vertex.Vertex[V]]E{}
		outgoing[e.GetDestination()] = e.Value()
		d.incomingToOutgoing[e.GetOrigin()] = outgoing
	} else {
		outgoing[e.GetDestination()] = e.Value()
	}
	if _, ok := d.uniqueVerticies[e.GetOrigin()]; !ok {
		d.uniqueVerticies[e.GetOrigin()] = struct{}{}
		d.vertices = append(d.vertices, e.GetOrigin())
	}
	if _, ok := d.uniqueVerticies[e.GetDestination()]; !ok {
		d.uniqueVerticies[e.GetDestination()] = struct{}{}
		d.vertices = append(d.vertices, e.GetDestination())
	}
}

// OutgoingVertices implements [graph.Graph].
func (d *dag[V, E]) OutgoingVertices(v vertex.Vertex[V]) func(yield func(vertex.Vertex[V]) bool) {
	outgoingVertex := d.incomingToOutgoing[v]

	return maps.Keys(outgoingVertex)
}

// Vertices implements [graph.Graph].
func (d *dag[V, E]) Vertices() []vertex.Vertex[V] {
	return d.vertices
}

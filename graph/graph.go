package graph

import (
	"github.com/gtantech/toposort/graph/edge"
	"github.com/gtantech/toposort/graph/vertex"
)

type Graph[V any, E any] interface {
	OutgoingEdges(vertex vertex.Vertex[V]) []edge.Edge[E, V]
	Vertices() []vertex.Vertex[V]
	AddVertex(v vertex.Vertex[V])
	AddEdge(e edge.Edge[E, V])
}

var _ Graph[string, string] = (*dag[string, string])(nil)

type dag[V any, E any] struct {
	vertices           []vertex.Vertex[V]
	incomingToOutgoing map[vertex.Vertex[V]]map[vertex.Vertex[V]]edge.Edge[E, V]
	uniqueVerticies    map[vertex.Vertex[V]]struct{}
}

func New[V any, E any]() Graph[V, E] {
	return &dag[V, E]{vertices: []vertex.Vertex[V]{}, incomingToOutgoing: make(map[vertex.Vertex[V]]map[vertex.Vertex[V]]edge.Edge[E, V]), uniqueVerticies: make(map[vertex.Vertex[V]]struct{})}
}

func (d *dag[V, E]) AddVertex(v vertex.Vertex[V]) {
	if _, ok := d.uniqueVerticies[v]; !ok {
		d.uniqueVerticies[v] = struct{}{}
		d.vertices = append(d.vertices, v)
	}
}

func (d *dag[V, E]) AddEdge(e edge.Edge[E, V]) {
	outgoing, ok := d.incomingToOutgoing[e.GetOrigin()]
	if !ok {
		outgoing := map[vertex.Vertex[V]]edge.Edge[E, V]{}
		outgoing[e.GetDestination()] = e
		d.incomingToOutgoing[e.GetOrigin()] = outgoing
	} else {
		outgoing[e.GetDestination()] = e
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

// OutgoingEdges implements [graph.Graph].
func (d *dag[V, E]) OutgoingEdges(v vertex.Vertex[V]) []edge.Edge[E, V] {
	outgoingVertex := d.incomingToOutgoing[v]

	result := make([]edge.Edge[E, V], 0, len(outgoingVertex))

	for key := range outgoingVertex {
		result = append(result, outgoingVertex[key])
	}

	return result //TODO change API to return iterator
}

// Vertices implements [graph.Graph].
func (d *dag[V, E]) Vertices() []vertex.Vertex[V] {
	return d.vertices
}

package graph

import (
	"maps"

	"github.com/gtantech/toposort/graph/vertex"
)

type Graph[V any, E any] interface {
	OutgoingVertices(vertex vertex.Vertex[V]) func(yield func(vertex.Vertex[V]) bool)
	Vertices() []vertex.Vertex[V]
	AddVertex(v vertex.Vertex[V])
	AddEdge(value E, origin vertex.Vertex[V], destination vertex.Vertex[V])
	GetEdgeValue(origin vertex.Vertex[V], destination vertex.Vertex[V]) E
	RemoveEdge(origin vertex.Vertex[V], destination vertex.Vertex[V])
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

func (d *dag[V, E]) RemoveEdge(origin vertex.Vertex[V], destination vertex.Vertex[V]) {
	outgoing, ok := d.incomingToOutgoing[origin]
	if ok {
		//origin vertex exists
		delete(outgoing, destination)
	}
}

func (d *dag[V, E]) GetEdgeValue(origin vertex.Vertex[V], destination vertex.Vertex[V]) E {
	outgoing, ok := d.incomingToOutgoing[origin]
	if ok {
		//origin vertex exists
		return outgoing[destination]
	}
	var zero E
	return zero
}

func (d *dag[V, E]) AddEdge(value E, origin vertex.Vertex[V], destination vertex.Vertex[V]) {
	outgoing, ok := d.incomingToOutgoing[origin]
	if !ok {
		outgoing := map[vertex.Vertex[V]]E{}
		outgoing[destination] = value
		d.incomingToOutgoing[origin] = outgoing
	} else {
		outgoing[destination] = value
	}
	if _, ok := d.uniqueVerticies[origin]; !ok {
		d.uniqueVerticies[origin] = struct{}{}
		d.vertices = append(d.vertices, origin)
	}
	if _, ok := d.uniqueVerticies[destination]; !ok {
		d.uniqueVerticies[destination] = struct{}{}
		d.vertices = append(d.vertices, destination)
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

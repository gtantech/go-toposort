package graph

import (
	"maps"

	"github.com/gtantech/toposort/graph/vertex"
)

type Graph[V any, E any] interface {
	OutgoingVertices(vertex vertex.Vertex[V]) func(yield func(vertex.Vertex[V]) bool)
	Vertices() func(yield func(vertex.Vertex[V]) bool)
	AddVertex(v vertex.Vertex[V])
	AddEdge(value E, origin vertex.Vertex[V], destination vertex.Vertex[V])
	GetEdgeValue(origin vertex.Vertex[V], destination vertex.Vertex[V]) (E, bool)
	RemoveEdge(origin vertex.Vertex[V], destination vertex.Vertex[V])
}

var _ Graph[string, string] = (*dag[string, string])(nil)

type dag[V any, E any] struct {
	edgeValues      map[vertex.Vertex[V]]map[vertex.Vertex[V]]E
	uniqueVerticies map[vertex.Vertex[V]]struct{}
}

func New[V any, E any]() Graph[V, E] {
	return &dag[V, E]{edgeValues: make(map[vertex.Vertex[V]]map[vertex.Vertex[V]]E), uniqueVerticies: make(map[vertex.Vertex[V]]struct{})}
}

func (d *dag[V, E]) AddVertex(v vertex.Vertex[V]) {
	if _, ok := d.uniqueVerticies[v]; !ok {
		d.uniqueVerticies[v] = struct{}{}
	}
}

func (d *dag[V, E]) RemoveEdge(origin vertex.Vertex[V], destination vertex.Vertex[V]) {
	outgoing, ok := d.edgeValues[origin]
	if ok {
		//origin vertex exists
		delete(outgoing, destination)
	}
}

func (d *dag[V, E]) GetEdgeValue(origin vertex.Vertex[V], destination vertex.Vertex[V]) (E, bool) {
	destinations, ok := d.edgeValues[origin]
	var zero E
	if !ok {
		return zero, false
	}

	//origin vertex exists
	value, ok := destinations[destination]

	if !ok {
		return zero, false
	}
	return value, true
}

func (d *dag[V, E]) AddEdge(value E, origin vertex.Vertex[V], destination vertex.Vertex[V]) {
	destinations, ok := d.edgeValues[origin]
	if !ok {
		outgoing := map[vertex.Vertex[V]]E{}
		outgoing[destination] = value
		d.edgeValues[origin] = outgoing
	} else {
		destinations[destination] = value
	}
	if _, ok := d.uniqueVerticies[origin]; !ok {
		d.uniqueVerticies[origin] = struct{}{}
	}
	if _, ok := d.uniqueVerticies[destination]; !ok {
		d.uniqueVerticies[destination] = struct{}{}
	}
}

// OutgoingVertices implements [graph.Graph].
func (d *dag[V, E]) OutgoingVertices(v vertex.Vertex[V]) func(yield func(vertex.Vertex[V]) bool) {
	destinations := d.edgeValues[v]

	return maps.Keys(destinations)
}

// Vertices implements [graph.Graph].
func (d *dag[V, E]) Vertices() func(yield func(vertex.Vertex[V]) bool) {
	return maps.Keys(d.uniqueVerticies)
}

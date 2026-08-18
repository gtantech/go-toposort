package edge

import (
	"github.com/gtantech/go-toposort/graph/vertex"
)

type Edge[E any, V any] interface {
	GetDestination() vertex.Vertex[V]
	GetOrigin() vertex.Vertex[V]
}

var _ Edge[string, string] = (*edge[string, string])(nil) //ensures edge implements Edge at compile time

type edge[E any, V any] struct {
	value       E
	origin      vertex.Vertex[V]
	destination vertex.Vertex[V]
}

func New[E any, V any](value E, origin vertex.Vertex[V], destination vertex.Vertex[V]) Edge[E, V] {
	return &edge[E, V]{value: value, origin: origin, destination: destination}
}

// GetDestination implements [graph.Edge].
func (e *edge[E, V]) GetDestination() vertex.Vertex[V] {
	return e.destination
}

// GetOrigin implements [Edge].
func (e *edge[E, V]) GetOrigin() vertex.Vertex[V] {
	return e.origin
}

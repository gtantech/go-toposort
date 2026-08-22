package toposort

import (
	"fmt"

	"github.com/gtantech/toposort/graph/edge"
	"github.com/gtantech/toposort/graph/vertex"
)

type CycleDetectedError[V any, E any] struct {
	Edge        edge.Edge[E, V]
	Origin      vertex.Vertex[V]
	Destination vertex.Vertex[V]
}

func (e *CycleDetectedError[V, E]) Error() string {
	return fmt.Sprintf("encountered cycle in graph for edge: %v from %v to %v", e.Edge.Value(), e.Origin.Value(), e.Destination.Value())
}

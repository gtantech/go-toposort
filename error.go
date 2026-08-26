package toposort

import (
	"fmt"

	"github.com/gtantech/toposort/graph/vertex"
)

type CycleDetectedError[V any, E any] struct {
	EdgeValue   E
	Origin      vertex.Vertex[V]
	Destination vertex.Vertex[V]
}

func (e *CycleDetectedError[V, E]) Error() string {
	return fmt.Sprintf("encountered cycle in graph for edge: %v from %v to %v", e.EdgeValue, e.Origin.Value(), e.Destination.Value())
}

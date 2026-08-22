package toposort

import (
	"fmt"

	"github.com/gtantech/toposort/graph/edge"
)

type CycleDetectedError[V any, E any] struct {
	Edge edge.Edge[E, V]
}

func (e *CycleDetectedError[V, E]) Error() string {
	return fmt.Sprintf("encountered cycle in graph for edge: %v from %v to %v", e.Edge.Value(), e.Edge.GetOrigin().Value(), e.Edge.GetDestination().Value())
}

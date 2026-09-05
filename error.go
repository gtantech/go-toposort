package toposort

import (
	"fmt"
)

type CycleDetectedError[V comparable, E any] struct {
	EdgeValue   E
	Origin      V
	Destination V
}

func (e *CycleDetectedError[V, E]) Error() string {
	return fmt.Sprintf("encountered cycle in graph for edge: %v from %v to %v", e.EdgeValue, e.Origin, e.Destination)
}

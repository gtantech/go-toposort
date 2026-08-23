package toposort

import (
	"errors"
	"fmt"

	"github.com/gtantech/toposort/graph"
	"github.com/gtantech/toposort/graph/vertex"
	"github.com/gtantech/toposort/stack"
)

func dfsTopo[V any, E any](g graph.Graph[V, E], v vertex.Vertex[V], s stack.Stack[vertex.Vertex[V]]) ([]vertex.Vertex[V], error) {
	order := []vertex.Vertex[V]{}
	s.Push(v)
	for !s.IsEmpty() {
		top := s.Peek()
		outgoingEdges := g.OutgoingEdges(top)
		foundUnexploredEdge := false
		for _, e := range outgoingEdges {
			if foundUnexploredEdge {
				break
			}
			opposite := e.GetDestination()
			if !opposite.IsVisited() {
				//discovery edge
				foundUnexploredEdge = true
				if err := s.Push(opposite); err != nil {
					if _, ok := errors.AsType[*stack.DuplicateValuesError[vertex.Vertex[V]]](err); ok {
						return nil, &CycleDetectedError[V, E]{Edge: e}
					}
					panic(fmt.Sprintf("unhandled error: %v", err))
				}
			}
		}
		if !foundUnexploredEdge {
			//all edges explored
			order = append([]vertex.Vertex[V]{top}, order...)
			visited := s.Pop()
			visited.SetVisited()
		}
	}
	return order, nil
}

func TopologicalSort[V any, E any](g graph.Graph[V, E]) ([]vertex.Vertex[V], error) {
	order := []vertex.Vertex[V]{}
	for _, v := range g.Vertices() {
		if !v.IsVisited() {
			suborder, err := dfsTopo(g, v, stack.New[vertex.Vertex[V]]())
			if err != nil {
				return nil, err
			}
			order = append(suborder, order...)
		}
	}
	return order, nil
}

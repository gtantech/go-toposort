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
		foundUnexploredEdge := false
		for opposite := range g.OutgoingVertices(top) {
			if foundUnexploredEdge {
				break
			}
			if opposite.IsVisited() {
				continue
			}
			//discovery edge
			foundUnexploredEdge = true
			pushErr := s.Push(opposite)
			if pushErr == nil {
				continue
			}
			//check error type
			if _, ok := errors.AsType[*stack.DuplicateValuesError[vertex.Vertex[V]]](pushErr); ok {
				edgeValue, ok := g.GetEdgeValue(top, opposite)
				if !ok {
					panic(fmt.Sprintf("failed to get edge value between %v and %v", top, opposite))
				}
				return nil, &CycleDetectedError[V, E]{EdgeValue: edgeValue, Origin: top, Destination: opposite}
			}
			panic(fmt.Sprintf("unhandled error for stack push: %v", pushErr))
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

package toposort

import (
	"errors"
	"fmt"

	"github.com/gtantech/toposort/graph"
	"github.com/gtantech/toposort/graph/vertex"
	"github.com/gtantech/toposort/stack"
)

func dfsTopo[V any, E any](g graph.Graph[V, E], v vertex.Vertex[V], s stack.Stack[vertex.Vertex[V]], isVisited map[vertex.Vertex[V]]bool) ([]vertex.Vertex[V], error) {
	order := []vertex.Vertex[V]{}
	s.Push(v)
	for !s.IsEmpty() {
		top := s.Peek()
		foundUnexploredEdge := false
		for opposite := range g.OutgoingVertices(top) {
			if foundUnexploredEdge {
				break
			}
			if isVisited[opposite] {
				continue
			}
			//discovery edge
			foundUnexploredEdge = true
			pushErr := s.Push(opposite)
			if pushErr == nil {
				continue
			}
			//check error type
			var e *stack.DuplicateValuesError[vertex.Vertex[V]]
			if errors.As(pushErr, &e) {
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
			isVisited[visited] = true
		}
	}
	return order, nil
}

func TopologicalSort[V any, E any](g graph.Graph[V, E]) ([]vertex.Vertex[V], error) {
	order := []vertex.Vertex[V]{}
	isVisited := make(map[vertex.Vertex[V]]bool)
	for v := range g.Vertices() {
		isVisited[v] = false
	}
	for v := range g.Vertices() {
		if !isVisited[v] {
			suborder, err := dfsTopo(g, v, stack.New[vertex.Vertex[V]](), isVisited)
			if err != nil {
				return nil, err
			}
			order = append(suborder, order...)
		}
	}
	return order, nil
}

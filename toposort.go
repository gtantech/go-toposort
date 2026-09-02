package toposort

import (
	"errors"
	"fmt"

	"github.com/gtantech/toposort/graph"
	"github.com/gtantech/toposort/stack"
)

func dfsTopo[V comparable, E any](g graph.Graph[V, E], v V, s stack.Stack[V], isVisited map[V]bool) ([]V, error) {
	order := []V{}
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
			var e *stack.DuplicateValuesError[V]
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
			order = append([]V{top}, order...)
			visited := s.Pop()
			isVisited[visited] = true
		}
	}
	return order, nil
}

func TopologicalSort[V comparable, E any](g graph.Graph[V, E]) ([]V, error) {
	order := []V{}
	isVisited := make(map[V]bool)
	for v := range g.Vertices() {
		isVisited[v] = false
	}
	for v := range g.Vertices() {
		if !isVisited[v] {
			suborder, err := dfsTopo(g, v, stack.New[V](), isVisited)
			if err != nil {
				return nil, err
			}
			order = append(suborder, order...)
		}
	}
	return order, nil
}

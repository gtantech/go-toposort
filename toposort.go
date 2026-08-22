package toposort

import (
	"fmt"

	"github.com/gtantech/toposort/graph"
	"github.com/gtantech/toposort/graph/vertex"
	"github.com/gtantech/toposort/stack"
)

func dfsTopo[V any, E any](g graph.Graph[V, E], v vertex.Vertex[V]) ([]vertex.Vertex[V], error) {
	s := stack.New[vertex.Vertex[V]]()
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
					return nil, fmt.Errorf("encountered cycle in graph for edge: %v from %v to %v", e.Value(), top.Value(), opposite.Value())
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
			suborder, err := dfsTopo(g, v)
			if err != nil {
				return nil, err
			}
			order = append(suborder, order...)
		}
	}
	return order, nil
}

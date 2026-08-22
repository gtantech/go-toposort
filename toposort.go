package toposort

import (
	"github.com/gtantech/go-container/stack"
	"github.com/gtantech/toposort/graph"
	"github.com/gtantech/toposort/graph/vertex"
)

func dfsTopo[V any, E any](g graph.Graph[V, E], v vertex.Vertex[V]) []vertex.Vertex[V] {
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
				s.Push(opposite)
			}
		}
		if !foundUnexploredEdge {
			//all edges explored
			order = append([]vertex.Vertex[V]{top}, order...)
			visited := s.Pop()
			visited.SetVisited()
		}
	}
	return order
}

func TopologicalSort[V any, E any](g graph.Graph[V, E]) []vertex.Vertex[V] {
	order := []vertex.Vertex[V]{}
	for _, v := range g.Vertices() {
		if !v.IsVisited() {
			suborder := dfsTopo(g, v)
			order = append(suborder, order...)
		}
	}
	return order
}

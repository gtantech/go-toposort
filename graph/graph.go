package graph

import (
	"github.com/gtantech/go-toposort/graph/edge"
	"github.com/gtantech/go-toposort/graph/vertex"
)

type Graph[V any, E any] interface {
	OutgoingEdges(vertex vertex.Vertex[V]) []edge.Edge[E, V]
	Vertices() []vertex.Vertex[V]
	AddVertex(v vertex.Vertex[V])
	AddEdge(e edge.Edge[E, V])
}

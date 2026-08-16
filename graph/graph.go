package graph

type Edge[E any, V any] interface {
	GetDestination() Vertex[V]
}

type Vertex[V any] interface {
	IsVisited() bool
	SetVisited()
}

type Graph[V any, E any] interface {
	OutgoingEdges(vertex Vertex[V]) []Edge[E, V]
	Vertices() []Vertex[V]
}

package vertex

type Vertex[V any] interface {
	IsVisited() bool
	SetVisited()
	Value() V
}

var _ Vertex[int] = (*vertex[int])(nil) //ensures vertex implements Vertex at compile time
type vertex[V any] struct {
	value     V
	isVisited bool
}

func New[V any](value V) Vertex[V] {
	return &vertex[V]{value: value, isVisited: false}
}

// IsVisited implements [graph.Vertex].
func (v *vertex[V]) IsVisited() bool {
	return v.isVisited
}

// SetVisited implements [graph.Vertex].
func (v *vertex[V]) SetVisited() {
	v.isVisited = true
}

// Value implements [graph.Vertex].
func (v *vertex[V]) Value() V {
	return v.value
}

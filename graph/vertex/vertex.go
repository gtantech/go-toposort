package vertex

type Vertex[V any] interface {
	Value() V
}

var _ Vertex[int] = (*vertex[int])(nil) //ensures vertex implements Vertex at compile time
type vertex[V any] struct {
	value V
}

func New[V any](value V) Vertex[V] {
	return &vertex[V]{value: value}
}

// Value implements [graph.Vertex].
func (v *vertex[V]) Value() V {
	return v.value
}

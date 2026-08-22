package stack

import (
	"fmt"

	"github.com/gtantech/go-container/stack"
)

type uniqueStack[T comparable] struct {
	stack.Stack[T]
	uniqueCheck map[T]struct{}
}

type DuplicateValuesError[T comparable] struct {
	Value T
}

func (d *DuplicateValuesError[T]) Error() string {
	return fmt.Sprintf("pushed duplicate value: %v", d.Value)
}

func New[T comparable]() *uniqueStack[T] {
	return &uniqueStack[T]{
		Stack:       stack.New[T](),
		uniqueCheck: make(map[T]struct{}),
	}
}

// will throw an error when pushed a duplicate value
func (s *uniqueStack[T]) Push(value T) error {
	if _, ok := s.uniqueCheck[value]; ok {
		return &DuplicateValuesError[T]{Value: value}
	}
	s.Stack.Push(value)
	s.uniqueCheck[value] = struct{}{}
	return nil
}

func (s *uniqueStack[T]) Pop() T {
	popped := s.Stack.Pop()
	delete(s.uniqueCheck, popped)
	return popped
}

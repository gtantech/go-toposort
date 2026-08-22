package stack

import (
	"errors"
	"fmt"
	"testing"
)

func TestPushPop(t *testing.T) {
	s := New[int]()
	err := s.Push(1)
	if err != nil {
		t.Errorf("unexpected error:%v", err)
	}
	if got := s.Pop(); got != 1 {
		t.Errorf("got %v, want 1", got)
	}
}

func TestPushDuplicate(t *testing.T) {
	s := New[int]()
	s.Push(1)
	err := s.Push(1)
	if err == nil {
		t.Errorf("expected error")
	}
	var dupErr *DuplicateValuesError[int]
	if !errors.As(err, &dupErr) {
		t.Errorf("expected error type")
	}
	if got := dupErr.Value; got != 1 {
		t.Errorf("got %v, want 1", got)
	}
	if got := dupErr.Error(); got != fmt.Sprintf("pushed duplicate value: %v", dupErr.Value) {
		t.Errorf("unexpected error msg: got %v, want %v", got, fmt.Sprintf("pushed duplicate value: %v", dupErr.Value))
	}
	if got := s.Pop(); got != 1 {
		t.Errorf("got %v, want 1", got)
	}
}

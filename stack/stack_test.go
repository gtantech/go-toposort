package stack

import "testing"

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
	if got := s.Pop(); got != 1 {
		t.Errorf("got %v, want 1", got)
	}
}

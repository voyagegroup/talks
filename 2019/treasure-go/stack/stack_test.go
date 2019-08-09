package stack

import (
	"testing"
)

func TestPushAndPop(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	if r := s.Pop(); r != 1 {
		t.Fatalf("want 1, got %v", r)
	}
	if r := s.Pop(); r != nil {
		t.Fatalf("want nil, got %v", r)
	}
}

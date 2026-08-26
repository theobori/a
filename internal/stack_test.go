package internal

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewEmptyStack[int]()

	for i := range 100 {
		s.Push(i)
		if s.Len() != i+1 {
			t.Fatal()
		}
	}

	for i := 99; i >= 0; i-- {
		if s.Top() != i {
			t.Fatal()
		}

		v, err := s.Pop()
		if err != nil {
			t.Fatal(err)
		}

		if v != i {
			t.Fatal()
		}

		if s.Len() != i {
			t.Fatal()
		}
	}
}

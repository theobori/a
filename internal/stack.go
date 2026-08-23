package internal

import (
	"fmt"
)

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) Push(elem T) {
	s.arr = append(s.arr, elem)
}

func (s *Stack[T]) Top() T {
	return s.arr[len(s.arr)-1]
}

func (s *Stack[T]) Pop() (T, error) {
	var ans T

	size := s.Len()
	if size == 0 {
		return ans, fmt.Errorf("Empty stack.")
	}

	ans = s.arr[size-1]
	s.arr = s.arr[0 : size-1]

	return ans, nil
}

func (s *Stack[T]) Len() int {
	return len(s.arr)
}

func NewStack[T any](l []T) *Stack[T] {
	s := Stack[T]{
		arr: []T{},
	}

	for _, elem := range l {
		s.Push(elem)
	}

	return &s
}

func NewEmptyStack[T any]() *Stack[T] {
	return NewStack([]T{})
}

package main

import (
	"fmt"
)

type Stack[T any] interface {
	// Push adds an element to the top of the stack
	Push(e T)
	// Pop removes and returns the top element of the stack
	Pop() (T, error)
	// Len returns the number of elements in the stack
	Len() int
	// Peek returns the top element of the stack without removing it
	Peek() (T, error)
	// IsEmpty returns true if the stack is empty
	IsEmpty() bool
}

var errStackEmpty = fmt.Errorf("stack is empty")

type stackImpl[T any] struct {
	inner []T
}

func New[T any]() Stack[T] {
	return &stackImpl[T]{}
}

func (s *stackImpl[T]) Push(e T) {
	s.inner = append(s.inner, e)
}

func (s *stackImpl[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errStackEmpty
	}

	last := s.inner[s.Len()-1]
	s.inner = s.inner[:s.Len()-1]
	return last, nil
}

func (s *stackImpl[T]) Len() int {
	return len(s.inner)
}

func (s *stackImpl[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errStackEmpty
	}

	return s.inner[s.Len()-1], nil
}

func (s *stackImpl[T]) IsEmpty() bool {
	return s.Len() == 0
}

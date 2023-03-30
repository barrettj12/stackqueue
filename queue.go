package main

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
	Len() int
}

// Basic queue implementation using a slice
type sliceQueue[T any] struct {
	elems []T
}

var _ Queue[any] = &sliceQueue[any]{}

func newSliceQueue[T any]() Queue[T] {
	return &sliceQueue[T]{}
}

func (s *sliceQueue[T]) Enqueue(t T) {
	s.elems = append(s.elems, t)
}

func (s *sliceQueue[T]) Dequeue() T {
	popped := s.elems[0]
	s.elems = s.elems[1:]
	return popped
}

func (s *sliceQueue[T]) Len() int {
	return len(s.elems)
}

// Implementation of queue using two stacks
// We enqueue by pushing onto stack1. When we need to dequeue, pop everything
// off stack1 and push in reverse order onto stack2, except the last element,
// which is returned. If we want to start enqueueing again, move everything
// back to stack1.
type queueFromStacks[T any] struct {
	stack1, stack2 Stack[T]
}

var _ Queue[any] = &queueFromStacks[any]{}

func newQueueFromStacks[T any]() Queue[T] {
	return &queueFromStacks[T]{
		stack1: newSliceStack[T](),
		stack2: newSliceStack[T](),
	}
}

func (s *queueFromStacks[T]) Enqueue(t T) {
	for s.stack2.Len() > 0 {
		s.stack1.Push(s.stack2.Pop())
	}
	s.stack1.Push(t)
}

func (s *queueFromStacks[T]) Dequeue() T {
	for s.stack1.Len() > 1 {
		s.stack2.Push(s.stack1.Pop())
	}

	if s.stack1.Len() > 0 {
		return s.stack1.Pop()
	} else {
		return s.stack2.Pop()
	}
}

func (s *queueFromStacks[T]) Len() int {
	return s.stack1.Len() + s.stack2.Len()
}

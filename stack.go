package main

type Stack[T any] interface {
	Push(T)
	Pop() T
	Len() int
}

// Basic stack implementation using a slice
type sliceStack[T any] struct {
	elems []T
}

var _ Stack[any] = &sliceStack[any]{}

func newSliceStack[T any]() Stack[T] {
	return &sliceStack[T]{}
}

func (s *sliceStack[T]) Push(t T) {
	s.elems = append(s.elems, t)
}

func (s *sliceStack[T]) Pop() T {
	popped := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return popped
}

func (s *sliceStack[T]) Len() int {
	return len(s.elems)
}

// Implementation of stack using two queues
// At any time, one queue is empty, and the other queue contains all the
// elements. When we pop, we have to transfer everything to the other queue.
type stackFromQueues[T any] struct {
	queue1, queue2 Queue[T]
}

var _ Stack[any] = &stackFromQueues[any]{}

func newStackFromQueues[T any]() Stack[T] {
	return &stackFromQueues[T]{
		queue1: newSliceQueue[T](),
		queue2: newSliceQueue[T](),
	}
}

func (s *stackFromQueues[T]) Push(t T) {
	if s.queue2.Len() > 0 {
		s.queue2.Enqueue(t)
	} else {
		s.queue1.Enqueue(t)
	}
}

func (s *stackFromQueues[T]) Pop() T {
	if s.Len() == 0 {
		panic("stack is empty")
	}

	// Work out which queue is empty
	var from, to Queue[T]
	if s.queue1.Len() > 0 {
		from = s.queue1
		to = s.queue2
	} else {
		from = s.queue2
		to = s.queue1
	}

	for from.Len() > 1 {
		to.Enqueue(from.Dequeue())
	}

	// One element left in `from` - this is the one we want
	return from.Dequeue()
}

func (s *stackFromQueues[T]) Len() int {
	return s.queue1.Len() + s.queue2.Len()
}

package main

import "testing"

var stackImplementations = map[string]func() Stack[string]{
	"slice stack":            newSliceStack[string],
	"stack using two queues": newStackFromQueues[string],
}

func TestStack(t *testing.T) {
	for name, newStack := range stackImplementations {
		t.Run(name, func(t *testing.T) {
			stack := newStack()
			// []
			checkLen(t, stack, 0)

			stack.Push("a")
			// [a]
			checkLen(t, stack, 1)

			stack.Push("b")
			// [a b]
			checkLen(t, stack, 2)

			popAndCheck(t, stack, "b")
			// [a]
			checkLen(t, stack, 1)

			stack.Push("c")
			// [a c]
			checkLen(t, stack, 2)

			stack.Push("d")
			// [a c d]
			checkLen(t, stack, 3)

			popAndCheck(t, stack, "d")
			// [a c]
			checkLen(t, stack, 2)

			popAndCheck(t, stack, "c")
			// [a]
			checkLen(t, stack, 1)

			popAndCheck(t, stack, "a")
			// []
			checkLen(t, stack, 0)

			stack.Push("e")
			// [e]
			checkLen(t, stack, 1)

			popAndCheck(t, stack, "e")
			// [e]
			checkLen(t, stack, 0)
		})
	}
}

func popAndCheck(t *testing.T, stack Stack[string], expected string) {
	popped := stack.Pop()
	if popped != expected {
		t.Errorf("stack.Pop returned %q, expected %q", popped, expected)
	}
}

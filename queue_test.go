package main

import "testing"

var queueImplementations = map[string]func() Queue[string]{
	"slice queue":            newSliceQueue[string],
	"queue using two stacks": newQueueFromStacks[string],
}

func TestQueue(t *testing.T) {
	for name, newQueue := range queueImplementations {
		t.Run(name, func(t *testing.T) {
			queue := newQueue()
			// []
			checkLen(t, queue, 0)

			queue.Enqueue("a")
			// [a]
			checkLen(t, queue, 1)

			queue.Enqueue("b")
			// [a b]
			checkLen(t, queue, 2)

			dequeueAndCheck(t, queue, "a")
			// [b]
			checkLen(t, queue, 1)

			queue.Enqueue("c")
			// [b c]
			checkLen(t, queue, 2)

			queue.Enqueue("d")
			// [b c d]
			checkLen(t, queue, 3)

			dequeueAndCheck(t, queue, "b")
			// [c d]
			checkLen(t, queue, 2)

			dequeueAndCheck(t, queue, "c")
			// [d]
			checkLen(t, queue, 1)

			dequeueAndCheck(t, queue, "d")
			// []
			checkLen(t, queue, 0)

			queue.Enqueue("e")
			// [e]
			checkLen(t, queue, 1)

			dequeueAndCheck(t, queue, "e")
			// [e]
			checkLen(t, queue, 0)
		})
	}
}

func dequeueAndCheck(t *testing.T, queue Queue[string], expected string) {
	dequeued := queue.Dequeue()
	if dequeued != expected {
		t.Errorf("queue.Dequeue returned %q, expected %q", dequeued, expected)
	}
}

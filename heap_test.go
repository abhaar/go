package _go

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinIntHeap(t *testing.T) {
	h := NewHeap(MinHeap)
	heap.Init(h)

	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 1)
	heap.Push(h, 4)

	for h.Len() > 0 {
		assert.Equal(t, 1, heap.Pop(h))
		assert.Equal(t, 2, heap.Pop(h))
		assert.Equal(t, 3, heap.Pop(h))
		assert.Equal(t, 4, heap.Pop(h))
	}
}

func TestMaxIntHeap(t *testing.T) {
	h := NewHeap(MaxHeap)
	heap.Init(h)

	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 1)
	heap.Push(h, 4)

	for h.Len() > 0 {
		assert.Equal(t, 4, heap.Pop(h))
		assert.Equal(t, 3, heap.Pop(h))
		assert.Equal(t, 2, heap.Pop(h))
		assert.Equal(t, 1, heap.Pop(h))
	}
}

package _go

type IntHeap struct {
	elements []int
	t        HeapType
}

func NewHeap(k HeapType) *IntHeap {
	return &IntHeap{
		elements: make([]int, 0),
		t:        k,
	}
}

func (h IntHeap) Len() int {
	return len(h.elements)
}

func (h IntHeap) Less(i, j int) bool {
	if h.t == MaxHeap {
		return h.elements[i] > h.elements[j]
	}
	return h.elements[i] < h.elements[j]
}

func (h IntHeap) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *IntHeap) Push(x interface{}) {
	h.elements = append(h.elements, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := h.elements
	n := len(old)
	x := old[n-1]
	h.elements = old[:n-1]
	return x
}

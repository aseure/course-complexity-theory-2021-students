package tree

import "fmt"

type MaxHeap struct {
	nodes []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		nodes: nil,
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *MaxHeap) String() string {
	return fmt.Sprintf("%v", h.nodes)
}

func (h *MaxHeap) Size() int {
	return len(h.nodes)
}

func (h *MaxHeap) Insert(value int) {
	h.nodes = append(h.nodes, value)
	h.heapifyUp(len(h.nodes) - 1)
}

func (h *MaxHeap) PopMax() (int, bool) {
	if len(h.nodes) == 0 {
		return 0, false
	}
	max := h.nodes[0]
	h.nodes[0] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.heapifyDown(0)
	return max, true
}

func (h *MaxHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.nodes[parent] < h.nodes[i] {
			h.swap(parent, i)
		}
		i = parent
	}
}

func (h *MaxHeap) heapifyDown(i int) {
	for {
		leftChild, rightChild := 2*i+1, 2*i+2
		isLeftChildBigger := leftChild < len(h.nodes) && h.nodes[i] < h.nodes[leftChild]
		isRightChildBigger := rightChild < len(h.nodes) && h.nodes[i] < h.nodes[rightChild]

		if isLeftChildBigger && isRightChildBigger {
			if h.nodes[leftChild] < h.nodes[rightChild] {
				h.swap(i, rightChild)
				i = rightChild
			} else {
				h.swap(i, leftChild)
				i = leftChild
			}
		} else if isLeftChildBigger {
			h.swap(i, leftChild)
			i = leftChild
		} else if isRightChildBigger {
			h.swap(i, rightChild)
			i = rightChild
		} else {
			return
		}
	}
}

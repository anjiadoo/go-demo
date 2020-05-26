package sort

import (
	"container/heap"
	"testing"
)

func TestInsertHeap(t *testing.T) {
	slice := []int{9, 4, 2, 6, 8, 0, 3, 1, 7, 5}
	t.Log(slice)

	heap := NewHeap(10)
	for _, val := range slice {
		heap.insert(val)
	}
	t.Log(heap.arr)
}

func TestHeapSort(t *testing.T) {
	slice := []int{-1, 5, 4, 2, 6, 8, 0, 3, 1, 7, 9}
	t.Log(slice)

	HeapSort(slice)
	t.Log(slice)
}

//////////////////////////////////////////////////////////////////////
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func TestIntHeap(t *testing.T) {
	h := &IntHeap{5, 4, 2, 6, 8, 0, 3, 1, 7, 9}
	heap.Init(h)
	for h.Len() > 0 {
		t.Logf("%d ", heap.Pop(h))
	}
}

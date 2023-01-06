package main

// Example of using go's standard min-heap implementation as a max-heap
// by means of data-swizzling

// This example demonstrates an integer heap built using the heap interface.
import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap.Init(h)
	fmt.Printf("\n%v\n", h)

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	fmt.Printf("\n%v\n", h)
}

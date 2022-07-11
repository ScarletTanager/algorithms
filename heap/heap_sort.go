package heap

import "golang.org/x/exp/constraints"

// Sort implements heapsort.  The slice passed in is sorted in situ, and the
// function returns the sorted slice as well.
// Heapsort works by moving the current topmost node to the end - in a max
// heap, this node is always guaranteed to be the largest value - and then truncating
// the heap before re-maxifying from the top.  This brings the next largest
// value up to the topmost node, which is then swapped with the last node in the
// heap (this is why we truncate after every swap).
// Heapsort runs in Θ(nlogn).
func Sort[O constraints.Ordered](d []O) []O {
	h := NewMaxHeap(d)
	for i := h.length; i >= 2; i-- {
		h.swap(1, i)
		h.truncate()
		h.Maxify(1)
	}
	return h.Data()
}

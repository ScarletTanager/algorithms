package heap

import (
	"math"

	"golang.org/x/exp/constraints"
)

type HeapType int

const (
	HEAPTYPE_MAX HeapType = iota
	HEAPTYPE_MIN
)

// type Numeric interface {
// 	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
// }

// Heap is, well, a heap.  It can store any of the golang built-in numeric types
// (anything which implements the <>== operators)
type Heap[O constraints.Ordered] struct {
	// Type tells us if this is a max or min heap
	Type   HeapType
	length int
	data   []O
}

// NewMaxHeap creates a new Max Heap (largest value in topmost node) from the slice
func NewMaxHeap[O constraints.Ordered](d []O) *Heap[O] {
	// Figure out how large to make our array
	l := 0
	for exp := 0; l < len(d); exp++ {
		l += int(math.Pow(2, float64(exp)))
	}

	dHeader := make([]O, 1)
	h := &Heap[O]{
		Type:   HEAPTYPE_MAX,
		length: l,
		data:   append(dHeader, d...),
	}

	var start int = h.length / 2
	for i := start; i >= 1; i = i - 1 {
		h.maxify(i)
	}

	return h
}

// Data returns the underlying slice used for heap storage
func (h *Heap[O]) Data() []O {
	return h.data[1:]
}

// maxify runs heap maxification from the specified node
func (h *Heap[O]) maxify(i int) {
	var cur, m int
	m = h.max((i * 2), ((i * 2) + 1))

	for cur = i; m != 0 && h.data[m] > h.data[cur]; {
		h.swap(cur, m)
		cur = m
		m = h.max((cur * 2), ((cur * 2) + 1))
	}
}

// Utility method which verifies that the position is within the actual heap.
func (h *Heap[O]) isValidPosition(i int) bool {
	return i > 0 && i < len(h.data)
}

// swap exchanges the values at the specified positions with one another
func (h *Heap[O]) swap(i, j int) {
	var tmp O = h.data[i]
	h.data[i] = h.data[j]
	h.data[j] = tmp
}

func (h *Heap[O]) truncate() {
	h.data = h.data[:len(h.data)-1]
}

// max returns the position containing the greater value
// Return value of 0 means neither position is valid
func (h *Heap[O]) max(i, j int) int {
	return h.minOrMax(i, j, HEAPTYPE_MAX)
}

// min returns the position containing the smaller value, i if both equal.
// Return value of 0 means neither position is valid
func (h *Heap[O]) min(i, j int) int {
	return h.minOrMax(i, j, HEAPTYPE_MIN)
}

// Implemention of min() and max()
func (h *Heap[O]) minOrMax(i, j int, op HeapType) int {
	var valI, valJ O

	if h.isValidPosition(i) {
		valI = h.data[i]
		if h.isValidPosition(j) {
			// Both indices are valid, so compare vals
			valJ = h.data[j]

			if op == HEAPTYPE_MAX {
				if valI >= valJ {
					return i
				}
			} else if op == HEAPTYPE_MIN {
				if valI <= valJ {
					return i
				}
			}
			return j
		} else {
			// only i valid
			return i
		}
	} else if h.isValidPosition(j) {
		// only j valid
		return j
	}

	// neither valid
	return 0
}

func parent(pos int) int {
	return pos >> 1
}

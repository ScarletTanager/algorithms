package heap

import "golang.org/x/exp/constraints"

func Sort[O constraints.Ordered](d []O) []O {
	h := NewMaxHeap(d)
	for i := h.length; i >= 2; i-- {
		h.swap(1, i)
		h.size--
		h.Maxify(1)
	}
	return h.Data()
}

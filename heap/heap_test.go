package heap_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/heap"
)

var _ = Describe("Heap", func() {
	var (
		data []int
		h    *heap.Heap[int]
	)

	BeforeEach(func() {
		data = []int{
			4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
		}
		//            4
		//        1       3
		//    2      16 9    10
		// 14    8  7
	})

	It("Creates a heap", func() {
		h = heap.NewMaxHeap(data)
		hd := h.Data()
		Expect(hd).To(Equal([]int{5, 5, 5}))
	})

	Describe("A slice sized `correctly`", func() {
		BeforeEach(func() {
			data = []int{
				4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 22, 5, 7, 13, 6,
			}
		})

		It("Creates a heap", func() {
			h = heap.NewMaxHeap(data)
			Expect(h.Data()).To(Equal([]int{5, 5, 5}))
		})
	})
})

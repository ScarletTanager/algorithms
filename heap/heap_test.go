package heap_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/heap"
)

var _ = Describe("Heap", func() {
	var (
		data, expected []int
		h              *heap.Heap[int]
	)

	BeforeEach(func() {
		data = []int{
			4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
		}
		//            4
		//        1       3
		//    2      16 9    10
		// 14    8  7
		expected = []int{
			16, 14, 10, 8, 7, 9, 3, 2, 4, 1,
		}
		//              16
		//          14       10
		//     8      7   9  3
		//   2   4  1
	})

	It("Creates a heap", func() {
		h = heap.NewMaxHeap(data)
		hd := h.Data()
		for i := 0; i < len(expected); i++ {
			Expect(hd[i]).To(Equal(expected[i]))
		}
	})

	Describe("A slice sized `correctly`", func() {
		BeforeEach(func() {
			data = []int{
				4, 1, 3, 2, 16, 9, 10, 14, 8, 7, 22, 5, 7, 13, 6,
			}
			//                  4
			//           1              3
			//      2        16      9        10
			//  14     8  7     22 5   7   13     6
			expected = []int{
				22, 16, 13, 14, 7, 9, 10, 2, 8, 4, 1, 5, 7, 3, 6,
			}
			//                  22
			//          16              13
			//     14        7      9          10
			//   2    8    4   1  5    7    3      6
			for _, i := range []int{0, 1, 2, 3, 4, 7, 9, 10, 13} {
				Expect(data[i]).NotTo(Equal(expected[i]))
			}
		})

		It("Creates a heap", func() {
			h = heap.NewMaxHeap(data)
			hd := h.Data()
			for i := 0; i < len(expected); i++ {
				Expect(hd[i]).To(Equal(expected[i]))
			}
		})

		It("Heapifies the original slice in situ", func() {
			heap.NewMaxHeap(data)
			for i := 0; i < len(expected); i++ {
				Expect(data[i]).To(Equal(expected[i]))
			}
		})
	})
})

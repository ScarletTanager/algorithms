package heap_test

import (
	"github.com/ScarletTanager/algorithms/heap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HeapSort", func() {
	var (
		data, expected []int
	)

	BeforeEach(func() {
		data = []int{9, 13, 1, 27, 5, 4, 3}
		expected = []int{1, 3, 4, 5, 9, 13, 27}
		for i := 0; i < len(expected); i++ {
			Expect(data[i]).NotTo(Equal(expected[i]))
		}
	})

	It("Sorts the slice in ascending order", func() {
		heap.Sort(data)
		for i := 0; i < len(expected); i++ {
			Expect(data[i]).To(Equal(expected[i]))
		}
	})
})

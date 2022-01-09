package sort_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/sort"
)

var _ = Describe("InsertionSort", func() {
	var (
		unsortedInts []int
	)

	BeforeEach(func() {
		unsortedInts = []int{9, 13, 1, 5, 27, 4, 3}
	})

	It("Sorts the values", func() {
		sorted := sort.InsertionSort(unsortedInts)
		Expect(sorted).To(Equal([]int{1, 3, 4, 5, 9, 13, 27}))
	})

	Describe("InsertionSortNonIncreasing", func() {
		It("Sorts the values", func() {
			sorted := sort.InsertionSortNonIncreasing(unsortedInts)
			Expect(sorted).To(Equal([]int{27, 13, 9, 5, 4, 3, 1}))
		})
	})
})

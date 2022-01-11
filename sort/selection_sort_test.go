package sort_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/sort"
)

var _ = Describe("SelectionSort", func() {
	var (
		unsortedInts []int
	)

	BeforeEach(func() {
		unsortedInts = []int{9, 13, 1, 5, 27, 4, 3}
	})

	It("Sorts the values", func() {
		sorted := sort.SelectionSort(unsortedInts)
		Expect(sorted).To(Equal([]int{1, 3, 4, 5, 9, 13, 27}))
	})
})

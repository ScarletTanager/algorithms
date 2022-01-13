package sort_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/sort"
)

var _ = Describe("MergeSort", func() {
	Describe("Merge", func() {
		var (
			a, b []int
		)

		BeforeEach(func() {
			a = []int{1, 4, 6, 7, 9}
			b = []int{2, 3, 5, 8, 10}
		})

		It("Merges the two arrays into a sorted and merged result", func() {
			sorted := sort.Merge(a, b)
			Expect(sorted).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
		})
	})

	Describe("MergeSort", func() {
		var (
			unsorted []int
		)

		BeforeEach(func() {
			unsorted = []int{10, 7, 1, 4, 2, 3, 8, 6, 5, 9}
		})

		It("Sorts the slice", func() {
			sorted := sort.MergeSort(unsorted)
			Expect(sorted).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
		})

		Context("When the number of elements is odd", func() {
			BeforeEach(func() {
				unsorted = []int{10, 7, 1, 4, 11, 2, 3, 8, 6, 5, 9}
			})

			It("Sorts the slice", func() {
				sorted := sort.MergeSort(unsorted)
				Expect(sorted).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
			})
		})
	})
})

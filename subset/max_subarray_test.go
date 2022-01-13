package subset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/subset"
)

var _ = Describe("MaxSubarray", func() {
	var (
		a []int
	)

	Describe("MaxCrossingSubarray", func() {

		BeforeEach(func() {
			a = []int{5, 1, -13, 8, -3, 4, -6, 0, 0, 7}
		})

		It("Finds the maximum subarray crossing the midpoint", func() {
			s, l, r := subset.MaxCrossingSubarray(a)
			Expect(s).To(Equal(10))
			Expect(l).To(Equal(3))
			Expect(r).To(Equal(9))
		})
	})

	Describe("MaxSubarray", func() {
		BeforeEach(func() {
			a = []int{9, -20, -100, 13, 12, 0, 4, -76, 243, 17, 16, -343}
		})

		It("Finds the maximum subarray", func() {
			s, l, r := subset.MaxSubarray(a)
			Expect(l).To(Equal(8))
			Expect(r).To(Equal(10))
			Expect(s).To(Equal(276))
		})
	})
})
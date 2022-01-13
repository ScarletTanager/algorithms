package subset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/subset"
)

var _ = Describe("MaxSubarray", func() {
	Describe("MaxCrossingSubarray", func() {
		var (
			a []int
		)

		BeforeEach(func() {
			a = []int{5, 1, -13, 8, -3, 4, -6, 0, 0, 7}
		})

		It("Finds the maximum subarray crossing the midpoint", func() {
			m := subset.MaxCrossingSubarray(a)
			Expect(m).To(Equal([]int{8, -3, 4, -6, 0, 0, 7}))
		})
	})
})

package subset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/subset"
)

var _ = Describe("Interval", func() {
	var (
		temps []float64
	)

	BeforeEach(func() {
		// Daily average temperature (C) for May 2018 in New Delhi, India
		temps = []float64{
			32.13,
			29.98,
			29.72,
			31.39,
			35.03,
			32.67,
			34.06,
			29.38,
			29.04,
			32.74,
			35.67,
			35.57,
			30.99,
			30.57,
			32.53,
			35.56,
			35.67,
			33.76,
			34.48,
			36.04,
			37.49,
			37.68,
			38.00,
			36.71,
			37.56,
			38.02,
			38.39,
			36.94,
			37.37,
			35.04,
			37.83,
		}
	})

	Describe("MaxAverageInterval", func() {

		It("Finds the maximum average interval", func() {
			m, l, r := subset.MaxAverageInterval(temps, 3)
			Expect(m).To(Equal(37.99))
			Expect(l).To(Equal(24))
			Expect(r).To(Equal(26))
		})

		Context("When the list has the same number of elements as specified for the interval", func() {
			It("Returns the average for the entire list", func() {
				m, l, r := subset.MaxAverageInterval(temps[10:13], 3)
				Expect(m).To(Equal((temps[10] + temps[11] + temps[12]) / float64(3)))
				Expect(l).To(Equal(0))
				Expect(r).To(Equal(2))
			})
		})

		Context("When the max interval is at the end of the list", func() {
			It("Returns the average for the correct interval", func() {
				m, l, r := subset.MaxAverageInterval(temps[10:27], 3)
				Expect(l).To(Equal(14))
				Expect(r).To(Equal(16))
				Expect(m).To(Equal((temps[24] + temps[25] + temps[26]) / float64(3)))
			})
		})
	})

	Describe("LongestIntervalAverageExceeds", func() {
		It("Finds the longest interval with an average > 38 C", func() {
			l, r := subset.LongestIntervalExceeds(temps, float64(38.00))
			Expect(l).To(Equal(25))
			Expect(r).To(Equal(26))
		})

		It("Finds the longest interval with an average > 35 C", func() {
			l, r := subset.LongestIntervalExceeds(temps, float64(35.00))
			Expect(l).To(Equal(19))
			Expect(r).To(Equal(30))
		})
	})

})

package matrix_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/matrix"
)

var _ = Describe("Matrix", func() {
	var (
		rows, cols int
	)

	Describe("NewMatrix", func() {
		BeforeEach(func() {
			rows = 5
			cols = 6
		})

		It("Returns a matrix of the specified size", func() {
			m := matrix.NewMatrix(rows, cols)
			Expect([][]int(m)).To(HaveLen(rows))
			for _, row := range m {
				Expect(row).To(HaveLen(cols))
			}
		})
	})

	Describe("Multiply", func() {
		var (
			mA, mB                     matrix.Matrix
			rowsA, colsA, rowsB, colsB int
		)

		BeforeEach(func() {
			rowsA = 3
			colsA = 4
			rowsB = 4
			colsB = 3

			mA = matrix.NewMatrix(rowsA, colsA)
			mB = matrix.NewMatrix(rowsB, colsB)

			for rn := range mA {
				for cn := range mA[rn] {
					mA[rn][cn] = cn + 1 // 1...n, n is the number of columns
				}
			}

			for rn := range mB {
				for cn := range mB[rn] {
					mB[rn][cn] = rn + 1 // mB should be the transverse of mA
				}
			}
		})

		It("Returns a correct matrix product", func() {
			mC, _ := matrix.Multiply(mA, mB)
			Expect(len(mC)).To(Equal(rowsA))
			for _, row := range mC {
				Expect(row).To(HaveLen(colsB))

				for _, val := range row {
					Expect(val).To(Equal(30))
				}
			}
		})
	})
})

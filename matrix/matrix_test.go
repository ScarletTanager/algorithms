package matrix_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/matrix"
)

var _ = Describe("Matrix", func() {
	var (
		rows, cols int
		mA, mB     matrix.Matrix
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

	Describe("Negate", func() {
		BeforeEach(func() {
			rows = 4
			cols = 4
			mA = matrix.NewMatrix(rows, cols)
			for rn := range mA {
				for cn := range mA[rn] {
					mA[rn][cn] = rn * cn
				}
			}
		})

		It("Negates each element of the matrix", func() {
			mANeg := matrix.Negate(mA)
			Expect(len(mANeg)).To(Equal(rows))
			for rn, row := range mANeg {
				Expect(len(row)).To(Equal(cols))
				for cn := range mANeg[rn] {
					Expect(mANeg[rn][cn]).To(Equal(cn * rn * (-1)))
				}
			}
		})
	})

	Describe("Add", func() {
		BeforeEach(func() {
			rows = 4
			cols = 5

			mA = matrix.NewMatrix(rows, cols)
			mB = matrix.NewMatrix(rows, cols)

			for rn := 0; rn < rows; rn++ {
				for cn := 0; cn < cols; cn++ {
					mA[rn][cn] = rn + cn
					mB[rn][cn] = 8 - (rn + cn)
				}
			}
		})

		It("Adds the two matrices correctly", func() {
			mC, err := matrix.Add(mA, mB)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(mC)).To(Equal(4))
			for _, row := range mC {
				Expect(row).To(HaveLen(cols))

				for _, c := range row {
					Expect(c).To(Equal(8))
				}
			}
		})
	})

	Describe("Subtract", func() {
		BeforeEach(func() {
			rows = 4
			cols = 4
			mA = matrix.NewMatrix(rows, cols)
			mB = matrix.NewMatrix(rows, cols)

			for rn := 0; rn < rows; rn++ {
				for cn := 0; cn < cols; cn++ {
					mA[rn][cn] = rn * cn
					mB[rn][cn] = rn + cn
				}
			}
		})

		It("Subtracts the two matrices correctly", func() {
			mC, err := matrix.Subtract(mA, mB)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(mC)).To(Equal(rows))
			for rn, row := range mC {
				Expect(len(row)).To(Equal(cols))

				for cn, v := range row {
					Expect(v).To(Equal((rn * cn) + ((rn + cn) * (-1))))
				}
			}
		})
	})

	Describe("Multiply", func() {
		var (
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

package matrix

import (
	"errors"
)

// The top index specifies the row, the nested index the column.
// This is counter-intuitive if you are used to thinking in terms
// of Cartesian coordinate systems, since this is akin to (y, x).
type Matrix [][]int

// NewMatrix returns a zero-valued matrix of the specified size.
func NewMatrix(rows, cols int) Matrix {
	m := make([][]int, rows)
	for i := range m {
		m[i] = make([]int, cols)
	}

	return m
}

// Negate returns the negative of the specified matrix A,
// which is a matrix -A in which each element has been multiplied
// by -1.
func Negate(mA Matrix) Matrix {
	mAneg := NewMatrix(len(mA), len(mA[0]))

	for rn := range mAneg {
		for cn := range mAneg[rn] {
			mAneg[rn][cn] = (-1) * mA[rn][cn]
		}
	}

	return mAneg
}

// Add adds two r x c matrices, errors if dimensions are not the same.
func Add(mA, mB Matrix) (Matrix, error) {
	if len(mA) != len(mB) || len(mA[0]) != len(mB[0]) {
		return nil, errors.New("Cannot add differently sized matrices")
	}

	mC := NewMatrix(len(mA), len(mA[0]))

	for rn := range mA {
		for cn := range mA[rn] {
			mC[rn][cn] = mA[rn][cn] + mB[rn][cn]
		}
	}

	return mC, nil
}

// Subtract subtracts matrix mB from matrix mA.  This is implemented by
// adding mA and -mB at the cost of a constant multiplier, but we gain
// reuse and code simplicity.
func Subtract(mA, mB Matrix) (Matrix, error) {
	return Add(mA, Negate(mB))
}

// Multiply implements simple brute-force matrix multiplication
// of two compatible matrices.  Incompatibility results in an
// error being returned.
func Multiply(mA, mB Matrix) (Matrix, error) {
	// Check if compatible (cols in mA == rows in mB)
	if len(mA[0]) != len(mB) {
		return nil, errors.New("Matrices are not compatible")
	}

	rowCount := len(mA)
	colCount := len(mB[0])

	mC := NewMatrix(rowCount, colCount)

	for rn := range mC {
		for cn := range mC[rn] {
			for k := range mA[0] {
				mC[rn][cn] += mA[rn][k] * mB[k][cn]
			}
		}
	}

	return mC, nil
}

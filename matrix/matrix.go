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

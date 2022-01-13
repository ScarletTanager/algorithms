package subset

import (
	"fmt"
	"math/bits"
)

func minValueInt() int {
	if bits.UintSize == 32 {
		return -2147483648
	}

	// Let's assume 64 bits if it ain't 32
	return -9223372036854775808
}

func MaxCrossingSubarray(a []int) []int {
	var (
		leftBound, rightBound, leftSum, rightSum, sum, midpoint int
	)

	leftSum = minValueInt()
	rightSum = leftSum

	midpoint = len(a) / 2
	fmt.Printf("Midpoint: %d\n", midpoint)
	for idx := midpoint - 1; idx >= 0; idx-- {
		sum += a[idx]
		if sum > leftSum {
			leftSum = sum
			leftBound = idx
		}
	}

	sum = 0
	for idr, val := range a[midpoint:] {
		sum += val
		if sum > rightSum {
			rightSum = sum
			rightBound = idr
		}
	}

	return a[leftBound : midpoint+rightBound+1]
}

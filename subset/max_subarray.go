package subset

import (
	"math/bits"
)

func minValueInt() int {
	if bits.UintSize == 32 {
		return -2147483648
	}

	// Let's assume 64 bits if it ain't 32
	return -9223372036854775808
}

// MaxCrossingSubarray finds the maximum subarray which includes
// (crosses) the midpoint of the list.
// MaxCrossingSubarray runs in Θ(n).
func MaxCrossingSubarray(a []int) (sum, leftBound, rightBound int) {
	var (
		leftSum, rightSum, midpoint int
	)

	leftSum = minValueInt()
	rightSum = leftSum

	midpoint = len(a) / 2
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

	sum = leftSum + rightSum
	rightBound += midpoint

	return
}

// MaxSubarray implements a recursive divide-and-conquer algorithm
// for computing the maximum subarray of a given list.
// Changes would be required to support 0-length or nil
// slices.
// MaxSubarray runs in Θ(nlogn).
func MaxSubarray(a []int) (sum, leftBound, rightBound int) {
	if len(a) == 1 {
		return a[0], 0, 0
	}

	var (
		midpoint                      int
		leftSum, leftLow, leftHigh    int
		rightSum, rightLow, rightHigh int
		crossSum, crossLow, crossHigh int
	)

	midpoint = len(a) / 2

	leftSum, leftLow, leftHigh = MaxSubarray(a[:midpoint])
	rightSum, rightLow, rightHigh = MaxSubarray(a[midpoint:])
	crossSum, crossLow, crossHigh = MaxCrossingSubarray(a)

	if leftSum >= rightSum && leftSum >= crossSum {
		sum = leftSum
		leftBound = leftLow
		rightBound = leftHigh
	} else if rightSum >= leftSum && rightSum >= crossSum {
		sum = rightSum
		leftBound = rightLow + midpoint
		rightBound = rightHigh + midpoint
	} else {
		sum = crossSum
		leftBound = crossLow
		rightBound = crossHigh
	}

	return
}

// MaxSubarrayLinear implements finding the maximum subarray
// in linear time (using Kadane's algorithm).
// MaxSubarrayLinear runs in Θ(n).
func MaxSubarrayLinear(a []int) (max, leftBound, rightBound int) {
	var (
		currentSum, currentLeftBound, currentRightBound int
	)

	if len(a) == 0 {
		return 0, 0, 0
	}

	max = minValueInt()
	currentSum = max

	// for idx, val := range a {
	// 	currentSum += val
	//
	// 	if currentSum > max {
	// 		if val > currentSum {
	// 			max = val
	// 			currentSum = val
	// 			leftBound = idx
	// 		} else {
	// 			max = currentSum
	// 		}
	// 		rightBound = idx
	// 	}
	// }
	for idx, val := range a {
		// Is our current sum at or below zero?
		if currentSum <= 0 {
			// If the value is greater than the current sum, start a new subarray
			// Remember that we only get here if currentSum is <= 0
			if val > currentSum {
				currentSum = val
				currentLeftBound = idx
				currentRightBound = idx
			}
		} else {
			// Extend the current subarray
			currentSum += val
			currentRightBound = idx
		}

		if currentSum > max {
			max = currentSum
			leftBound = currentLeftBound
			rightBound = currentRightBound
		}
	}

	return
}

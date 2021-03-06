package subset

// MaxAverageInterval is a linear-time algoritm for finding the interval
// (subarray/slice) of the specified length with the maximum average
// value within the specified set of values.
// This is another implementation of Kadane's algorithm.
// MaxAverageInterval runs in Θ(n).
func MaxAverageInterval(a []float64, intervalLen int) (maxAverage float64,
	leftBound, rightBound int) {
	var (
		currentSum, maxSum float64
	)

	if len(a) == 0 || intervalLen <= 0 {
		return 0, 0, 0
	}

	maxSum = float64(-9223372036854775808)

	for idx, val := range a[:len(a)-intervalLen+1] {
		currentSum = val
		var j int
		for j = idx + 1; j < idx+intervalLen; j++ {
			currentSum += a[j]
		}

		if currentSum > maxSum {
			maxSum = currentSum
			leftBound = idx
			rightBound = j - 1
		}

	}
	maxAverage = maxSum / float64(intervalLen)

	return
}

func LongestIntervalExceeds(a []float64, threshold float64) (leftBound, rightBound int) {
	// Convert every value into a deviation
	devs := make([]float64, len(a))
	for i, v := range a {
		devs[i] = v - threshold
	}

	var (
		currentSum, maxSum                  float64
		currentLeftBound, currentRightBound int
	)

	maxSum = float64(-9223372036854775808)

	// Now we just run Kadane...
	for idx, val := range devs {
		if currentSum <= 0 {
			currentSum = val
			currentLeftBound = idx
		} else {
			currentSum += val
			currentRightBound = idx
		}

		if currentSum > maxSum {
			maxSum = currentSum
			leftBound = currentLeftBound
			rightBound = currentRightBound
		}
	}

	return
}

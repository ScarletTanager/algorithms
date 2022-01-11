package sort

// Insertion Sort - a well-known sorting algorithm
// Runs in Θ(n^2)

// InsertionSort performs a basic non-decreasinginsertion sort on
//the slice of integers, returning the sorted slice.  The original slice
// is sorted in situ.
func InsertionSort(a []int) []int {
	// Loop through the slice, starting with the second element.
	for index, val := range a[1:] {
		// This is just because the range operator always starts with the 0 index.
		// This would be unnecessary if we indexed directly into the slice
		// starting with element at position 1.
		index += 1
		inner := index - 1
		for inner >= 0 && a[inner] > val {
			// Shift everything to the right
			a[inner+1] = a[inner]
			inner -= 1 // We could use -- here, of course
		}

		// The for loop exits when we've found where our value goes
		a[inner+1] = val
	}

	return a
}

// InsertionSortNonIncreasing is identical to InsertionSort, except that
// it sorts in non-increasing order (as opposed to non-decreasing).
func InsertionSortNonIncreasing(a []int) []int {
	// Start with second-to-last element
	for index := len(a) - 2; index >= 0; index-- {
		val := a[index]
		inner := index + 1
		for inner < len(a) && a[inner] > val {
			// Shift values greater than `val` to the left
			a[inner-1] = a[inner]
			inner += 1
		}

		a[inner-1] = val
	}

	return a
}

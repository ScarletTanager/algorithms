package sort

// Selection Sort
// Runs in Θ(n^2)

func SelectionSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		pos := i + 1
		smallest := a[pos]
		for j := i + 1; j < len(a); j++ {
			if a[j] < smallest {
				smallest = a[j]
				pos = j
			}
		}
		if smallest < a[i] {
			a[pos] = a[i]
			a[i] = smallest
		}
	}
	return a
}

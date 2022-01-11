package sort

// Merge merges two sorted subarrays into a single merged/sorted
// array.  If either a or b is not already sorted, the behavior is
// undefined.
// Merge runs in Θ(n) where n is the total number of elements,
// i.e. len(a) + len(b)
func Merge(a, b []int) []int {
	totLen := len(a) + len(b)
	merged := make([]int, totLen)

	ida := 0
	idb := 0
	for idxMerged, _ := range merged {
		if (ida < len(a)) && (idb < len(b)) {
			if a[ida] <= b[idb] {
				merged[idxMerged] = a[ida]
				ida += 1
			} else {
				merged[idxMerged] = b[idb]
				idb += 1
			}
		} else if ida < len(a) {
			merged[idxMerged] = a[ida]
			ida += 1
		} else {
			merged[idxMerged] = b[idb]
			idb += 1
		}
	}
	return merged
}

// MergeSort is a reasonably efficient algorithm for sorting using divide/conquer
// via recursion.
// MergeSort runs in Θ(nlogn)
func MergeSort(a []int) []int {
	if len(a) > 1 {
		var midpoint int
		midpoint = len(a) / 2
		a1 := MergeSort(a[:midpoint])
		a2 := MergeSort(a[midpoint:])
		return Merge(a1, a2)
	}
	return a
}

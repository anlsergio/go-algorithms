package sort

// QuickSort is an implementation of the quicksort algorithm.
func QuickSort(iterable []int) []int {
	// items below length of 2 are already ordered
	// so simply return them.
	if len(iterable) < 2 {
		return iterable
	}

	// pivotIndex could be any index inside the slice range.
	pivotIndex := 0
	pivot := iterable[pivotIndex]
	var smaller, bigger []int

	// iterate over all items except for the pivot
	for _, item := range iterable[pivotIndex+1:] {
		if item <= pivot {
			smaller = append(smaller, item)
			continue
		}
		bigger = append(bigger, item)
	}

	smallerSorted := QuickSort(smaller)
	biggerSorted := QuickSort(bigger)

	sorted := append(smallerSorted, pivot)
	sorted = append(sorted, biggerSorted...)

	return sorted
}

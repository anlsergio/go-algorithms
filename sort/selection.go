package sort

// SelectionSort is an implementation of the selection sort algorithm.
func SelectionSort(iterable []int) []int {
	var sortedIterable []int

	for range iterable {
		smallerIndex := getSmallerIndex(iterable)
		sortedIterable = append(sortedIterable, iterable[smallerIndex])
		// iterable gets a new slice without the corresponding item at
		// smallerIndex by appending everything that is around it
		// since Go lang has no built-in "pop" method.
		iterable = append(iterable[:smallerIndex], iterable[smallerIndex+1:]...)
	}

	return sortedIterable
}

func getSmallerIndex(iterable []int) int {
	candidateIndex := 0
	candidate := iterable[candidateIndex]

	for i, item := range iterable {
		if item < candidate {
			candidate = item
			candidateIndex = i
		}
	}

	return candidateIndex
}

package search

// Binary is an implementation of a binary search
// which takes in an ordered iterable, an item to be found
// and returns an index to the corresponding position
// of the item if there's any.
//
// - (index == -1) means that the item does not exist.
//
// - Execution cost: O(log n)
func Binary(iterable []int, item int) (index int) {
	start := 0
	end := len(iterable) - 1

	for start <= end {
		middle := (start + end) / 2
		guess := iterable[middle]

		if guess == item {
			return middle
		}

		if guess < item {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}

	return -1
}

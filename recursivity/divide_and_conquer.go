package recursivity

// Sum takes in a slice of numbers and returns the sum of all items.
func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	// break down the problem into smaller pieces so that
	// the array eventually is reduced to 1 element being
	// sum up with another number.
	return numbers[0] + Sum(numbers[1:])
}

// CountItems counts the total number of items in items.
// Similarly to the built-in "len" but taking a
// recursive-divide-and-conquer approach.
func CountItems(items []int) int {
	if len(items) == 0 {
		return 0
	}

	return 1 + CountItems(items[1:])
}

// BiggestNumber returns the biggest number in numbers.
func BiggestNumber(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	firstItem := numbers[0]
	if len(numbers) == 1 {
		return firstItem
	}

	candidate := BiggestNumber(numbers[1:])

	if firstItem > candidate {
		return firstItem
	}

	return candidate
}

// BinarySearch is an implementation of a binary search
// which takes in an ordered iterable, an item to be found
// and returns true if the item is found in the iterable.
func BinarySearch(iterable []int, item int) (found bool) {
	if len(iterable) == 0 {
		return false
	}

	if len(iterable) == 1 && iterable[0] == item {
		return true
	}

	end := len(iterable) - 1
	middle := end / 2
	guess := iterable[middle]

	if item == guess {
		return true
	}

	if guess < item {
		return BinarySearch(iterable[middle+1:], item)
	}

	if guess > item {
		return BinarySearch(iterable[:middle], item)
	}

	return false
}

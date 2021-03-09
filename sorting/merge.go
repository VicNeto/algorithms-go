package sorting

// MergeSort func
func MergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}

	middle := int(len(array) / 2)
	left := MergeSort(array[:middle])
	right := MergeSort(array[middle:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	var result []int
	rightIndex, leftIndex := 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	for leftIndex < len(left) {
		result = append(result, left[leftIndex])
		leftIndex++
	}
	for rightIndex < len(right) {
		result = append(result, right[rightIndex])
		rightIndex++
	}

	return result
}

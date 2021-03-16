package divide_conquer

// MergeCountInv func
func MergeCountInv(array []int) ([]int, int) {
	if len(array) <= 1 {
		return array, 0
	}

	middle := int(len(array) / 2)
	left, left_count := MergeCountInv(array[:middle])
	right, right_count := MergeCountInv(array[middle:])
	merged, count_inv := merge_and_count_split_inv(left, right)
	return merged, left_count + right_count + count_inv
}

func merge_and_count_split_inv(left []int, right []int) ([]int, int) {
	var result []int
	rightIndex, leftIndex := 0, 0
	split_inv := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
			split_inv += len(left) - leftIndex
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

	return result, split_inv
}

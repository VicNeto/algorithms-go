package divideConquer

// MergeCountInv func
func MergeCountInv(array []int) ([]int, int) {
	if len(array) <= 1 {
		return array, 0
	}

	middle := int(len(array) / 2)
	left, leftCount := MergeCountInv(array[:middle])
	right, rightCount := MergeCountInv(array[middle:])
	merged, countInv := mergeAndCountSplitInv(left, right)
	return merged, leftCount + rightCount + countInv
}

func mergeAndCountSplitInv(left []int, right []int) ([]int, int) {
	var result []int
	rightIndex, leftIndex := 0, 0
	splitInv := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
			// important step to calculate split inversions
			splitInv += len(left) - leftIndex
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

	return result, splitInv
}

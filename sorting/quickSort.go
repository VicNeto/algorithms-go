package sorting

func QuickSort(array *[]int, left int, right int) {
	if left >= right {
		return
	}
	pivot := choosePivot(left, right)
	(*array)[left], (*array)[pivot] = (*array)[pivot], (*array)[left]
	newPivot := partition(array, left, right)
	QuickSort(array, left, newPivot-1)
	QuickSort(array, newPivot+1, right)
}

func partition(array *[]int, left int, right int) int {
	i := left + 1
	for j := left + 1; j < right; j++ {
		if (*array)[j] < (*array)[left] {
			(*array)[j], (*array)[i] = (*array)[i], (*array)[j]
			i++
		}
	}
	(*array)[left], (*array)[i-1] = (*array)[i-1], (*array)[left]
	return i - 1
}

func choosePivot(left int, right int) int {
	// return int((left + right) / 2)
	return right - 1
}

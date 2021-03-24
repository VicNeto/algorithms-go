package sorting

import (
	"math/rand"
	"time"
)

func QuickSort(array *[]int, left int, right int) {
	if left >= right {
		return
	}
	pivot := choosePivot(left, right)
	(*array)[left], (*array)[pivot] = (*array)[pivot], (*array)[left]
	newPivot := partition(array, left, right)
	QuickSort(array, left, newPivot)
	QuickSort(array, newPivot+1, right)
}

func partition(array *[]int, left int, right int) int {
	arr := *array
	i := left + 1
	for j := left + 1; j < right; j++ {
		if arr[j] < arr[left] {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[left], arr[i-1] = arr[i-1], arr[left]
	return i - 1
}

func choosePivot(left int, right int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(right-left) + left
}

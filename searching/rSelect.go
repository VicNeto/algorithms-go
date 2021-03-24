package searching

import (
	"math/rand"
	"time"
)

func RSelect(array []int, itemPos int) int {
	if len(array) == 1 {
		return array[0]
	}
	pivot := choosePivot(len(array))
	array[0], array[pivot] = array[pivot], array[0]
	position := partition(&array)

	if position == itemPos {
		return array[position]
	} else if position > itemPos {
		return RSelect(array[:position], itemPos)
	}
	return RSelect(array[position+1:], itemPos-position-1)
}

func partition(array *[]int) int {
	arr := *array
	i := 1
	for j := 1; j < len(arr); j++ {
		if arr[j] < arr[0] {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[0], arr[i-1] = arr[i-1], arr[0]
	return i - 1
}

func choosePivot(length int) int {
	rand.Seed(time.Now().UnixNano())
	if length <= 0 {
		return 0
	}
	return rand.Intn(length)
}

package sorting

import (
	"algoritms-go/dataStructures"
	"container/heap"
)

func HeapSort(array []int) []int {
	h := &dataStructures.IntHeap{}
	heap.Init(h)

	for _, elem := range array {
		heap.Push(h, elem)
	}
	sorted := []int{}

	for i := 0; i < len(array); i++ {
		sorted = append(sorted, heap.Pop(h).(int))
	}

	return sorted
}

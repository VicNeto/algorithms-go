package sorting

import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := []int{8, 5, 6, 1, 4, 2, 7, 3, 9}
	QuickSort(&array, 0, len(array))

	if !sort.IntsAreSorted(array) {
		var sorted []int
		copy(sorted, array)
		sort.Ints(sorted)
		t.Error("For", array, "expected", sorted, "got", array)
	}
}

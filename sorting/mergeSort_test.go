package sorting

import (
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	array := []int{8, 5, 6, 1, 4, 2, 7, 3, 9}
	sorted := MergeSort(array)

	if !sort.IntsAreSorted(sorted) {
		sort.Ints(array)
		t.Error("For", array, "expected", array, "got", sorted)
	}
}

package main

import (
	conquer "algoritms-go/divide_conquer"
	"algoritms-go/sorting"
	"fmt"
)

func main() {
	a := sorting.MergeSort([]int{7, 6, 9, 11, 1, 2, 5, 2, 3, 6, 1, 38})
	_, inv := conquer.MergeCountInv([]int{1, 3, 5, 2, 4, 6})
	fmt.Println(a)
	fmt.Println(inv)
}

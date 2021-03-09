package main

import (
	"algoritms-go/sorting"
	"fmt"
)

func main() {
	a := sorting.MergeSort([]int{7, 6, 9, 11, 1, 2, 5, 2, 3, 6, 1, 38})
	fmt.Println(a)
}

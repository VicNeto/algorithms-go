package main

import (
	"algoritms-go/graphs"
	"fmt"
)

func main() {
	// a := sorting.MergeSort([]int{7, 6, 9, 11, 1, 2, 5, 2, 3, 6, 1, 38})
	// _, inv := conquer.MergeCountInv([]int{1, 3, 5, 2, 4, 6})
	// fmt.Println(a)
	// fmt.Println(inv)
	// b := []int{3, 2, 7, 1, 6, 4, 5}
	// sorting.QuickSort(&b, 0, len(b))
	// fmt.Println(b)
	// c := []int{4, 2, 7, 1, 6, 5, 3}
	// iPos := searching.RSelect(c, 5)
	// fmt.Println(iPos)
	// fmt.Println(c)
	gr := graphs.UndirectedGraph([]string{"a", "b", "c", "d"}, [][]string{{"a", "b"}, {"a", "c"}, {"b", "c"}, {"b", "d"}, {"c", "d"}})
	gr.String()
	w := graphs.BreadFirstSearch(&gr, "a")
	fmt.Printf("%v\n", w)
	ugr := graphs.UndirectedGraph([]string{"a", "b", "c", "d", "e", "f", "g", "h"},
		[][]string{{"a", "b"}, {"a", "c"}, {"b", "c"}, {"d", "e"}, {"d", "f"}, {"g", "h"}})
	cc, ncc := graphs.UCC(&ugr)
	ugr.String()
	fmt.Printf("%v\n%d\n", cc, ncc)
	dgr := graphs.CreateGraph([]string{"a", "b", "c", "d"}, [][]string{{"a", "b"}, {"a", "c"}, {"b", "d"}, {"c", "d"}}, true, true)
	dgr.String()
	topo := graphs.TopologicalSort(&dgr)
	fmt.Printf("%v\n", topo)
}

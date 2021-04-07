package graphs

import (
	"container/list"
	"math"
)

func BreadFirstSearch(g *Graph, source string) map[string]int {
	unexplored, lengths := initializeVertices(g.vertices)
	lengths[source] = 0
	unexplored[source] = false
	queue := list.New()
	queue.PushBack(g.vertices[source])
	var v *Vertex
	var element *list.Element

	for queue.Len() > 0 {
		element = queue.Front()
		v = element.Value.(*Vertex)
		queue.Remove(element)
		for dest := range g.edges[v].destinations {
			if unexplored[dest.identifier] {
				unexplored[dest.identifier] = false
				lengths[dest.identifier] = lengths[v.identifier] + 1
				_ = queue.PushBack(dest)
			}
		}
	}

	return lengths
}

func initializeVertices(vertices map[string]*Vertex) (map[string]bool, map[string]int) {
	unexp := make(map[string]bool)
	lengths := make(map[string]int)
	for key := range vertices {
		unexp[key] = true
		lengths[key] = int(math.NaN())
	}
	return unexp, lengths
}

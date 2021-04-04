package graphs

import (
	"math"
)

func BreadFirstSearch(g *Graph, source string) map[string]int {
	unexplored, lengths := initializeVertices(g.vertices)
	lengths[source] = 0
	unexplored[source] = false
	queue := []*Vertex{}
	queue = append(queue, g.vertices[source])
	var v *Vertex

	for len(queue) > 0 {
		v = queue[0]
		queue = queue[1:]
		for dest, _ := range g.edges[v].destinations {
			if unexplored[dest.identifier] {
				unexplored[dest.identifier] = false
				lengths[dest.identifier] = lengths[v.identifier] + 1
				queue = append(queue, dest)
			}
		}
	}

	return lengths
}

func initializeVertices(vertices map[string]*Vertex) (map[string]bool, map[string]int) {
	unexp := make(map[string]bool)
	lengths := make(map[string]int)
	for key, _ := range vertices {
		unexp[key] = true
		lengths[key] = int(math.NaN())
	}
	return unexp, lengths
}

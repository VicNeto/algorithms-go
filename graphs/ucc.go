package graphs

func UCC(g *Graph) (map[*Vertex]int, int) {
	unexplored, _ := initializeVertices(g.vertices)
	cc := make(map[*Vertex]int)
	numCC := 0
	queue := []*Vertex{}
	var v *Vertex

	for _, i := range g.vertices {
		if unexplored[i.identifier] {
			numCC++
			queue = append(queue, g.vertices[i.identifier])
			for len(queue) > 0 {
				v = queue[0]
				queue = queue[1:]
				cc[i] = numCC
				if g.edges[v].destinations != nil {
					for dest := range g.edges[v].destinations {
						if unexplored[dest.identifier] {
							unexplored[dest.identifier] = false
							cc[dest] = numCC
							queue = append(queue, dest)
						}
					}
				}
			}
		}
		queue = []*Vertex{}
	}
	return cc, numCC
}

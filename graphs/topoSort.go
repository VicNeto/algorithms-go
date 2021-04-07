package graphs

var curLabel int
var unexp map[string]bool
var fValue map[string]int

func TopologicalSort(g *Graph) map[string]int {
	curLabel = len(g.vertices)
	unexp = make(map[string]bool)
	fValue = make(map[string]int)
	for key := range g.vertices {
		unexp[key] = true
	}

	for k, v := range g.vertices {
		if unexp[k] {
			dfsTopo(k, v, g)
		}
	}
	return fValue
}

func dfsTopo(k string, v *Vertex, g *Graph) {
	unexp[k] = false
	if g.edges[v] != nil {
		for d := range g.edges[v].destinations {
			if unexp[d.identifier] {
				dfsTopo(d.identifier, d, g)
			}
		}
	}
	fValue[k] = curLabel
	curLabel--
}

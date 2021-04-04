package graphs

import (
	"fmt"
	"strconv"
)

// Vertex (Node) type
// the value could change for generics when they come to go
type Vertex struct {
	identifier string
}

type Edges struct {
	destinations map[string]uint
}

type Graph struct {
	vertices []*Vertex
	edges    map[string]*Edges
}

func (g *Graph) AddVertex(v string) {
	nV := Vertex{v}
	g.vertices = append(g.vertices, &nV)
}

func (g *Graph) AddEdge(s, d, value string) {
	if g.edges == nil {
		g.edges = make(map[string]*Edges)
	}
	if g.edges[s] == nil {
		g.edges[s] = &Edges{}
		g.edges[s].destinations = make(map[string]uint)
	}
	val, _ := strconv.ParseInt(value, 10, 64)
	g.edges[s].destinations[d] = uint(val)
}

func (g *Graph) String() {
	for _, v := range g.vertices {
		fmt.Printf("%v ->", v.identifier)
		if g.edges[v.identifier] != nil {
			for dest, e := range g.edges[v.identifier].destinations {
				fmt.Printf("%s (%d)", dest, e)
			}
		}
		fmt.Println()
	}
}

func CreateGraph(vertices []string, edges [][]string) Graph {
	var g Graph
	for _, v := range vertices {
		g.AddVertex(v)
	}
	for _, e := range edges {
		g.AddEdge(e[0], e[1], "1")
	}
	return g
}

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
	destinations map[*Vertex]int
}

type Graph struct {
	vertices map[string]*Vertex
	edges    map[*Vertex]*Edges
}

func (g *Graph) AddVertex(v string) {
	nV := Vertex{v}
	if g.vertices == nil {
		g.vertices = make(map[string]*Vertex)
	}
	g.vertices[v] = &nV
}

func (g *Graph) AddEdge(s, d *Vertex, value string) {
	if g.edges == nil {
		g.edges = make(map[*Vertex]*Edges)
	}
	if g.edges[s] == nil {
		g.edges[s] = &Edges{}
		g.edges[s].destinations = make(map[*Vertex]int)
	}
	val, _ := strconv.ParseInt(value, 10, 64)
	g.edges[s].destinations[d] = int(val)
}

func (g *Graph) String() {
	for _, v := range g.vertices {
		fmt.Printf("%v ->", v.identifier)
		if g.edges[v] != nil {
			for dest, e := range g.edges[v].destinations {
				fmt.Printf("%s (%d)", dest, e)
			}
		}
		fmt.Println()
	}
}

// CreateGraph return Graph
func CreateGraph(vertices []string, edges [][]string, directed bool) Graph {
	var g Graph
	for _, v := range vertices {
		g.AddVertex(v)
	}
	for _, e := range edges {
		if g.vertices[e[0]] != nil && g.vertices[e[1]] != nil {
			g.AddEdge(g.vertices[e[0]], g.vertices[e[1]], "1")
			if !directed {
				g.AddEdge(g.vertices[e[1]], g.vertices[e[0]], "1")
			}
		}
	}
	return g
}

// UndirectedGraph return Graph
func UndirectedGraph(vertices []string, edges [][]string) Graph {
	return CreateGraph(vertices, edges, false)
}

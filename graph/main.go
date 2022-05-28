package main

import "fmt"

// Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVerex adds a Vertex to the Graph
func (g *Graph) AddVerex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %d not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check errors
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Edge already exists (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}
	}
	return false
}

// Print will print the ajacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex: %v :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
}

func main() {
	test := &Graph{}
	for i := 0; i <= 5; i++ {
		test.AddVerex(i)
	}

	test.AddVerex(0)
	test.AddEdge(1, 2)
	test.AddEdge(1, 2)
	test.AddEdge(6, 2)
	test.AddEdge(3, 2)
	test.Print()
}

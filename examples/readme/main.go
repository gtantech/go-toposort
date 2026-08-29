package main

import (
	"fmt"

	"github.com/gtantech/toposort"
	"github.com/gtantech/toposort/graph"
	"github.com/gtantech/toposort/graph/vertex"
)

func main() {
	g := graph.New[string, int]()

	//Creating vertices for graph
	A := vertex.New("A")
	B := vertex.New("B")
	C := vertex.New("C")

	//link vertices via edges
	/*
		   |A|   |B|
		      \ /
		       |
			   v
		      |C|
	*/

	//add edges to graph
	g.AddEdge(0, A, C)
	g.AddEdge(1, B, C)

	//get topological sorted order
	order, _ := toposort.TopologicalSort(g)

	//print order
	for _, v := range order {
		fmt.Printf("|%v|,", v.Value())
	}
}

package main

import (
	"fmt"

	"github.com/gtantech/toposort/v2"
	"github.com/gtantech/toposort/v2/graph"
)

func main() {
	g := graph.New[string, int]()

	//Creating vertices for graph
	A := "A"
	B := "B"
	C := "C"

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
		fmt.Printf("|%v|,", v)
	}
}

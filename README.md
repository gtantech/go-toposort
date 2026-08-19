# toposort
A Go package that provides a generic implementation to topological sort a graph.

# Install
Install via `go get`. Note that Go 1.25 or newer is required.

```sh
# After: go mod init ...
go get -u github.com/gtantech/toposort
```

# Example

```go
package main

import (
	"fmt"

	"github.com/gtantech/toposort"
	"github.com/gtantech/toposort/graph"
	"github.com/gtantech/toposort/graph/edge"
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
	AC := edge.New(0, A, C)
	BC := edge.New(0, B, C)

	//add edges to graph
	g.AddEdge(AC)
	g.AddEdge(BC)

	//get topological sorted order
	order := toposort.TopologicalSort(g)

	//print order
	for _, v := range order {
		fmt.Printf("|%v|,", v.Value())
	}
}

```

## User Defined Structs

The `TopologicalSort` method accepts any structs that implement the interfaces defined in the graph package (graph, edge, vertex).
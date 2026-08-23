# Topological Sorting for Golang
Toposort is a Go package that provides a generic implementation to topological sort a graph.

[![CI Status](https://github.com/gtantech/toposort/actions/workflows/ci.yaml/badge.svg)](https://github.com/gtantech/toposort/actions/workflows/ci.yaml) [![codecov](https://codecov.io/gh/gtantech/toposort/graph/badge.svg)](https://codecov.io/gh/gtantech/toposort) [![Docs](https://godoc.org/github.com/gtantech/toposort?status.svg)](https://pkg.go.dev/github.com/gtantech/toposort?tab=doc)
## Table of Contents
- [What is Topological Sorting](#what-is-topological-sorting)
- [Why is Topological Sorting Useful](#why-is-topological-sorting-useful)
- [Install](#install)
- [Example](#example)
- [User Defined Structs](#user-defined-structs)
- [Error Handling](#error-handling)
- [License](#license)

## What Is Topological Sorting
From [Wikipedia](https://en.wikipedia.org/wiki/Topological_sorting), topological sorting of a [directed acyclic graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph) (DAG) is a linear ordering of its vertices such that when the graph is traversed in that order, every node is visited only when all preceeding nodes connected by a directed edge are visited first.

## Why Is Topological Sorting Useful
Topological sorting can help order dependencies such as in task scheduling/project management or course scheduling where a prerequisite course must be taken before taking advanced courses.

## Install
Install via `go get`. Note that Go 1.25 or newer is required.

```sh
# After: go mod init ...
go get -u github.com/gtantech/toposort
```

## Example

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
	order, _ := toposort.TopologicalSort(g)

	//print order
	for _, v := range order {
		fmt.Printf("|%v|,", v.Value())
	}
}

```

## User Defined Structs

The `TopologicalSort` function accepts any structs that implement the interfaces defined in the graph package (graph, edge, vertex).

## Error Handling
`TopologicalSort` features cycle detection and will return a `CycleDetectedError` when encountering a cycle within the DAG. Below is an error handling example, extending from the above example:

```go
//get topological sorted order
order, err := toposort.TopologicalSort(g)

if err != nil {
	if err, ok := errors.AsType[*toposort.CycleDetectedError[string, int]](err); ok {
		fmt.Printf("error: encountered cycle within graph at edge %v from %v to %v",
			err.Edge.Value(),
			err.Edge.GetOrigin().Value(),
			err.Edge.GetDestination().Value(),
		)
	}
}
```

## License

Licensed under [MIT License](./LICENSE)

## Thanks!

Thanks for reading and happy coding! Add a star to the project if you find it useful!
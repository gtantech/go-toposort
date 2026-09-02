package main

import (
	"fmt"

	"github.com/gtantech/toposort/v2"
	"github.com/gtantech/toposort/v2/graph"
)

type Actvity struct {
	Name     string
	Duration int
}

func NewActivity(name string, durationMinutes int) *Actvity {
	n := Actvity{Name: name, Duration: durationMinutes}
	return &n
}

type Project struct {
	graph.Graph[*Actvity, int]
	Name string
}

func NewProject(name string) *Project {
	n := Project{Graph: graph.New[*Actvity, int](), Name: name}
	return &n
}

func (p *Project) AddDepdendency(predecessor *Actvity, successor *Actvity) {
	p.Graph.AddEdge(predecessor.Duration, predecessor, successor)
}

func main() {
	p := NewProject("morning routine")

	//Creating vertices for graph
	makeCoffee := NewActivity("make coffee", 4)
	makeToast := NewActivity("make toast", 4)
	eatBreakfast := NewActivity("eat breakfast", 10)
	brushTeeth := NewActivity("brush teeth", 2)
	shower := NewActivity("shower", 10)
	getReadyForWork := NewActivity("get ready for work", 10)

	p.AddDepdendency(brushTeeth, eatBreakfast)
	p.AddDepdendency(makeCoffee, makeToast)
	p.AddDepdendency(makeCoffee, eatBreakfast)
	p.AddDepdendency(makeToast, eatBreakfast)
	p.AddDepdendency(eatBreakfast, shower)
	p.AddDepdendency(shower, getReadyForWork)

	//get topological sorted order
	order, _ := toposort.TopologicalSort(p)

	//print order
	fmt.Printf("%v:\n", p.Name)
	for i, v := range order {
		if i < len(order)-1 {
			fmt.Printf("|%v|,", v.Name)
		} else {
			fmt.Printf("|%v|", v.Name)
		}
	}
}

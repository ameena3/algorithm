package main

import (
	"fmt"

	datastructure "github.com/ameena3/interview_practice/data_structure"
)

func main() {
	g := datastructure.NewGraph()
	root := g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddNode("D")
	g.AddVertex("A", "B", 0)
	g.AddVertex("B", "C", 0)
	g.Print()
	if datastructure.BFS(root, "C") {
		fmt.Println("we found the node")
	} else {
		fmt.Println("No node found")
	}
}

// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"fmt"

	hp "container/heap"

	datastructure "github.com/ameena3/interview_practice/data_structure"
)

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	g := datastructure.NewGraph()
	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddNode("D")
	g.AddNode("E")
	g.AddNode("F")
	g.AddNode("G")
	g.AddNode("H")
	g.AddNode("I")

	g.AddVertex("A", "B", 5)
	g.AddVertex("A", "C", 3)
	g.AddVertex("A", "E", 2)

	g.AddVertex("B", "D", 2)

	g.AddVertex("C", "B", 1)
	g.AddVertex("C", "D", 1)

	g.AddVertex("D", "A", 1)
	g.AddVertex("D", "H", 1)
	g.AddVertex("D", "G", 2)

	g.AddVertex("E", "H", 4)
	g.AddVertex("E", "I", 7)
	g.AddVertex("E", "A", 1)

	g.AddVertex("F", "B", 3)
	g.AddVertex("F", "G", 1)

	g.AddVertex("G", "C", 3)
	g.AddVertex("G", "I", 2)

	g.AddVertex("H", "C", 2)
	g.AddVertex("H", "G", 2)

	Dishtras("A", "I", g)
}

func Dishtras(start, end string, g *datastructure.Graph) {
	var heap hp.Interface = datastructure.NewPQ()
	hp.Init(heap)
	visited := make(map[string]bool)
	hp.Push(heap, &datastructure.Item{
		Value:    start,
		Path:     make([]*datastructure.Node, 0),
		Priority: 0,
	})
	for heap.Len() != 0 {
		item := hp.Pop(heap).(*datastructure.Item)
		if item.Value == end {
			for _, p := range item.Path {
				fmt.Printf("%s -> ", p.Key)
			}
			fmt.Printf("%s\n", end)
			fmt.Printf("Weight %d\n", item.Priority)
			return
		}
		node := g.GetNode(item.Value)
		for _, itemNode := range node.Adjacents {
			if !visited[itemNode.Key] {
				hp.Push(heap, &datastructure.Item{
					Value:    itemNode.Key,
					Path:     append(item.Path, node),
					Priority: item.Priority + itemNode.WeightFromNode,
				})
			}
		}
		visited[node.Key] = true
	}
}

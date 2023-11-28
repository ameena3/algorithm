package datastructure

import (
	"errors"
	"fmt"
)

type (

	/*
		Graph	{
			Node1 -> []Node
			Node2 -> []Node
			...
		}
	*/
	Graph struct {
		Nodes []*Node
	}

	Node struct {
		Key            string
		Marked         bool
		WeightFromNode int
		Adjacents      []*Node
	}
)

func NewGraph() *Graph {
	return &Graph{
		Nodes: make([]*Node, 0),
	}
}

// Check if the node exists and add it if not present
func (g *Graph) AddNode(key string) *Node {
	if ok, n := contains(g.Nodes, key); !ok {
		newNode := &Node{Key: key}
		g.Nodes = append(g.Nodes, newNode)
		return newNode
	} else {
		return n
	}
}

// Check the to and the from node if either one of them do not exist
// return error or noop if the vertex exist otherwise add one
func (g *Graph) AddVertex(from, to string, weight int) error {
	containsFrom, nodeFrom := contains(g.Nodes, from)
	containsTo, _ := contains(g.Nodes, to)
	if !containsFrom || !containsTo {
		return errors.New("there was no node pair for the vertex to be added please add the nodes first")
	}
	for _, n := range nodeFrom.Adjacents {
		if n.Key == to {
			return nil
		}
	}
	nodeFrom.Adjacents = append(nodeFrom.Adjacents, &Node{
		Key:            to,
		Marked:         false,
		WeightFromNode: weight,
	})
	return nil
}

func (g *Graph) Print() {
	fmt.Println("********GRAPH*********")
	for _, n := range g.Nodes {
		fmt.Printf("%s : ", n.Key)
		for _, ad := range n.Adjacents {
			fmt.Printf("%s -> %d ->", ad.Key, ad.WeightFromNode)
		}
		fmt.Println()
	}
	fmt.Println("")
	fmt.Println("********GRAPH_END*********")
}

func (g *Graph) GetNode(key string) *Node {
	if ok, n := contains(g.Nodes, key); ok {
		return n
	}
	return nil
}

func contains(list []*Node, key string) (bool, *Node) {
	for i, n := range list {
		if n.Key == key {
			return true, list[i]
		}
	}
	return false, nil
}

// Uses a queue and we process until the queue is not empty

func BFS(root *Node, key string) bool {
	queue := make([]*Node, 0)
	var elem *Node
	queue = append(queue, root)
	root.Marked = true
	for len(queue) != 0 {
		elem, queue = Dequeue(queue)
		if elem.Key == key {
			return true
		}
		for _, ad := range elem.Adjacents {
			if !ad.Marked {
				ad.Marked = true
				queue = append(queue, ad)
			}
		}
	}
	return false
}

func Dequeue(list []*Node) (*Node, []*Node) {
	length := len(list)
	if length <= 1 {
		return list[0], make([]*Node, 0)
	}
	element := list[len(list)-1]
	list = list[:len(list)-2]
	return element, list
}

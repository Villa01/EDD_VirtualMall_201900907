package main

type Node struct {
	data string
	next *Node
	previous *Node

}

type Nodes []Node

// Constructor de Nodos para lista doblemente enlazada
func NewNode(value string) *Node{
	node := Node{value, nil,nil}
	return &node
}


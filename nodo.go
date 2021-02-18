package main

type Nodes []Node

// NewNode es un constructor de Nodos para lista doblemente enlazada
func NewNode(value Store) *Node {
	node := Node{value, nil, nil}
	return &node
}

package main

import (
	"fmt"
	"strings"
)

//AddAtTheBeggining Agrega el nodo al principio de la lista y lo vuelve la cabeza
func (list *DoublyLinkedList) AddAtTheBeggining(node *Node) {
	if list.lenght == 0 {
		list.head = node

	} else {
		secondNode := list.head
		list.head = node
		list.head.next = secondNode
		secondNode.previous = node
	}
	list.lenght++
}

// Append  Agrega un nodo al final
func (list *DoublyLinkedList) Append(newNode *Node) {
	if list.head == nil {
		list.head = newNode
		list.lenght++
		return
	}
	lastNode := list.GetLastNode()
	lastNode.next = newNode
	newNode.previous = lastNode

	list.lenght++
}

//AddAfter Agrega un nodo delante de un indice
func (list *DoublyLinkedList) AddAfter(index int, newNode *Node) string {
	var err string
	var previousNode *Node
	var nextNode *Node
	if list.isEmpty() {
		err = "vacia"
	} else if list.isInRange(index) {
		err = "El indice está fuera de rango"
	} else {
		previousNode, _ = list.GetNodeAt(index)
		if previousNode.next != nil {
			nextNode = previousNode.next
		}

		previousNode.next = newNode
		newNode.previous = previousNode

		if nextNode != nil {
			nextNode.previous = newNode
			newNode.next = nextNode
		} else {
			newNode.next = nil
		}

	}
	list.lenght++
	return err
}

// Retorna true si la lista esta vacia
func (list *DoublyLinkedList) isEmpty() bool {
	return list.lenght == 0
}

// ToString Retorna la lista en 2 strings recorrida de atras hacia adelante y de adelante hacia atras
func (list DoublyLinkedList) ToString() (string, string) {
	var temp = list.head
	var text string
	var reText string

	if list.isEmpty() {
		text = "vacia"
		reText = text
	} else {
		for i := list.lenght; i > 0; i-- {
			if temp != list.head {
				text += " <- " + temp.data.Name
			} else {
				text += temp.data.Name
			}

			temp = temp.next
		}

		temp = list.GetLastNode()

		for i := list.lenght; i > 0; i-- {

			reText += " -> " + temp.data.Name
			temp = temp.previous
		}
	}

	return text, reText
}

//GetLastNode Retorna el último nodo de la lista
func (list DoublyLinkedList) GetLastNode() *Node {
	var lastNode *Node
	temp := list.head
	for i := 1; i < list.lenght; i++ {
		temp = temp.next
	}

	lastNode = temp
	return lastNode
}

//GetNodeAt Retorna la referencia al nodo en el indice solicitado
func (list DoublyLinkedList) GetNodeAt(index int) (*Node, string) {

	var err string
	var temp *Node
	if list.isEmpty() {
		err = "La lista se encuenta vacia"
	} else if list.isInRange(index) {
		err = "El indice está fuera del rango"
	} else {
		temp = list.head
		for i := 0; i < index; i++ {
			temp = temp.next
		}
	}

	return temp, err
}

// Evualua si el indice esta dentro del rango de la lista

func (list DoublyLinkedList) isInRange(index int) bool {
	return index < 0 || index > list.lenght-1
}

//SortByAscii ordena de menor a mayor por el codigo ascii del nombre de la tienda
func (list *DoublyLinkedList) SortByAscii() {

	for i := 0; i < list.lenght; i++ {

	}
}

func GetAsciiValue(word string) int {
	sum := 0
	for i := 0; i < len(word); i++ {
		fmt.Println("La letra es : ", word[i])
		sum += int(word[i])
	}
	return sum
}

//GetGraphviz retorna una representacion en texto en formato dot
func (list DoublyLinkedList) GetGraphviz() string {
	var text string

	var letter string

	text += "\trankdir = \"TB\"\n"
	letter = string(list.head.data.Name[0])

	var tempNode1 *Node
	var tempNode2 *Node
	i := 0
	for i < list.lenght-1 || i == 0 {

		if list.lenght == 1 {
			tempNode1 := list.GetLastNode()
			nodeName := letter + fmt.Sprint(tempNode1.data.Rating) + fmt.Sprint(i+1) + strings.ReplaceAll(tempNode1.data.Name, " ", "")

			text += "\tnode [ shape= rect label=\"" + tempNode1.data.Name + "\" ] " + nodeName + ";\n"
		} else {
			tempNode1, _ = list.GetNodeAt(i)
			tempNode2 = tempNode1.next
			nodeName := letter + fmt.Sprint(tempNode1.data.Rating) + fmt.Sprint(i) + strings.ReplaceAll(tempNode1.data.Name, " ", "")
			nodeName2 := letter + fmt.Sprint(tempNode2.data.Rating) + fmt.Sprint(i+1) + strings.ReplaceAll(tempNode1.data.Name, " ", "")
			text += "\tnode [ shape= rect label=\"" + tempNode1.data.Name + "\" ] " + nodeName + ";\n"
			text += "\tnode [ shape= rect label=\"" + tempNode2.data.Name + "\" ] " + nodeName2 + ";\n"
			text += "\t" + nodeName + " -> " + nodeName2 + ";\n"
			text += "\t" + nodeName2 + " -> " + nodeName + ";\n"
		}

		i++
	}

	return text
}

// GetJSONNodes regresa los nodos de una lista en un formato compatible con json
func (list DoublyLinkedList) GetJSONNodes() SeveralJSONNodes {

	var Anodes []JSONNodes
	temp := list.head
	for i := 0; i < list.lenght; i++ {

		Nname := temp.data.Name
		Ndescription := temp.data.Description
		Ncontact := temp.data.Contact

		node := JSONNodes{
			Name:        Nname,
			Description: Ndescription,
			Contact:     Ncontact,
		}

		Anodes = append(Anodes, node)
		temp, _ = list.GetNodeAt(i)
	}
	data := SeveralJSONNodes{
		Nodes: Anodes,
	}
	return data
}

// DeleteNode elimina el nodo de la lista
func (list *DoublyLinkedList) DeleteNode(index int) {

	if list.isEmpty() {
		return
	}
	node, _ := list.GetNodeAt(index)
	fmt.Println(node)
	fmt.Println(node.data.Name)

	if list.lenght == 1 {
		list.head = nil
	} else if node == list.GetLastNode() { // Si es el ultimo al anterior se le apunta a nil
		node.previous.next = nil
	} else if node == list.head { // Si es el primero al siguiente se le apunta a nil
		temp := node.next
		list.head = temp
		node.next.previous = nil

	} else {
		node.previous.next = node.next
		node.next.previous = node.previous
	}

	list.lenght--

}

func (list *DoublyLinkedList) searchByContent(storeName string) *Node{
	var node *Node

	i := 0
	for i < list.lenght{
		temp, _ := list.GetNodeAt(i)
		if temp.data.Name == storeName {
			node = temp
		}
		i++
	}

	return node
}

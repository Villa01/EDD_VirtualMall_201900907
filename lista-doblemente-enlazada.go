package main

type DoublyLinkedList struct {
	head   *Node
	lenght int
}

// Crea una DoublyLinkedList vacia
func NewDoublyLinkedList() DoublyLinkedList {
	return DoublyLinkedList{nil, 0}
}

// Agrega el nodo al principio de la lista y lo vuelve la cabeza
func (list *DoublyLinkedList) addAtTheBeggining(node *Node) {
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

// Agrega un nodo al final
func (list *DoublyLinkedList) add(newNode *Node) {
	if list.isEmpty() {
		list.head = newNode
		return
	} else {

		lastNode := list.getLastNode()
		lastNode.next = newNode
		newNode.previous = lastNode

	}
	list.lenght++
}

// Agrega un nodo delante de un indice
func (list *DoublyLinkedList) addAfter(index int, newNode *Node) string {
	var err string
	var previousNode *Node
	var nextNode *Node
	if list.isEmpty() {
		err = "La lista se encuentra vacia"
	} else if list.isInRange(index) {
		err = "El indice está fuera de rango"
	} else {
		previousNode, _ = list.getNodeAt(index)
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
	return list.head == nil
}

// Retorna la lista en 2 strings recorrida de atras hacia adelante y de adelante hacia atras
func (list DoublyLinkedList) toString() (string, string) {

	var temp = list.head
	var text string
	var reText string
	for i := list.lenght; i > 0; i-- {
		if temp != list.head {
			text += " <- " + temp.data
		} else {
			text += temp.data
		}

		temp = temp.next
	}

	temp = list.getLastNode()

	for i := list.lenght; i > 0; i-- {

		reText += " -> " + temp.data
		temp = temp.previous
	}
	return text, reText
}

// Retorna el último nodo de la lista
func (list DoublyLinkedList) getLastNode() *Node {
	var lastNode *Node
	temp := list.head
	for i := 1; i < list.lenght; i++ {
		temp = temp.next
	}

	lastNode = temp
	return lastNode
}

// Retorna la referencia al nodo en el indice solicitado
func (list DoublyLinkedList) getNodeAt(index int) (*Node, string) {

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

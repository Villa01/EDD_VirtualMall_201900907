package main


type listaMeses struct {
	lenght int `json:"lenght"`
	head *nodoMes
}


type nodoMes struct {
	mes int `json:"mes"`
	data *matriz `json:"data"`
	next     *nodoMes `json:"next"`
	previous *nodoMes `json:"previous"`

}

func nuevaListaMeses() *listaMeses{
	return &listaMeses{0, nil}
}

func nuevoNodoMes (mes int, data *matriz) *nodoMes{
	return &nodoMes{mes, data, nil, nil}
}

func (list *listaMeses) Append(newNode *nodoMes) {
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

func (list *listaMeses) isEmpty() bool {
	return list.lenght == 0
}

//GetLastNode Retorna el último nodo de la lista
func (list listaMeses) GetLastNode() *nodoMes {
	var lastNode *nodoMes
	temp := list.head
	for i := 1; i < list.lenght; i++ {
		temp = temp.next
	}

	lastNode = temp
	return lastNode
}

//GetNodeAt Retorna la referencia al nodo en el indice solicitado
func (list listaMeses) GetNodeAt(index int) (*nodoMes, string) {

	var err string
	var temp *nodoMes
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

func (list listaMeses) isInRange(index int) bool {
	return index < 0 || index > list.lenght-1
}

func (list *listaMeses) searchByContent(mes int) *nodoMes{
	var node *nodoMes

	i := 0
	for i < list.lenght{
		temp, _ := list.GetNodeAt(i)
		if temp.mes == mes {
			node = temp
		}
		i++
	}

	return node
}

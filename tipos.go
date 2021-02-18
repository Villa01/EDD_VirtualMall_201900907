package main

type DoublyLinkedList struct {
	head   *Node
	lenght int
}

// NewDoublyLinkedList crea una DoublyLinkedList vacia
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, 0}
}

// Node es un nodo con una tienda dentro
type Node struct {
	data     Store
	next     *Node
	previous *Node
}

// Store es un tipo donde se puede almacenar toda la información de una tienda
type Store struct {
	Name        string `json:"Nombre"`
	Description string `json:"Descripcion"`
	Contact     string `json:"Contacto"`
	Rating      int8   `json:"Calificacion"`
}

// Department se refiere a los departamentos a los que pertenencen las tiendas
type Department struct {
	Name   string  `json:"Nombre"`
	Stores []Store `json:"Tiendas"`
}

// Datum es una matriz 3 x 3 que contiene indices y departamentos
type Datum struct {
	Index       string       `json:"Indice"`
	Departments []Department `json:"Departamentos"`
}

// Information es un tipo que contiene informacion de las tienas de un centro comercial
type Information struct {
	Data []Datum `json:"Datos"`
}

//IndexLetter es el tipo de dato de las filas de la matriz
type IndexLetter struct {
	Index       string
	Departments []DepartmentMatrix
}

//DepartmentMatrix es el tipo de dato de las columnas de la matriz
type DepartmentMatrix struct {
	Department
	name    string
	ratings [5]Rating
}

// Rating es la tercera dimension de la matriz
type Rating struct {
	number int
	lista  *DoublyLinkedList
}

// Matrix es una matriz llenada con informacion de un json
type Matrix struct {
	indexes []*IndexLetter
}

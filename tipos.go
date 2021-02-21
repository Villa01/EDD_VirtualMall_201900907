package main

// DoublyLinkedList es una lista doble enlazada
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
	Departments []DepartmentMatrix `json: "Departamentos`
}

//DepartmentMatrix es el tipo de dato de las columnas de la matriz
type DepartmentMatrix struct {
	Name    string    `json:"Departamento"`
	Ratings [5]Rating `json:"Calificaciones"`
}

// Rating es la tercera dimension de la matriz
type Rating struct {
	Number int `json: "Calificacion"`
	Lista  *DoublyLinkedList
}

// Matrix es una matriz llenada con informacion de un json
type Matrix struct {
	Indexes []*IndexLetter
}

// JSONNodes es una representacion de una store compatible con json
type JSONNodes struct {
	Name        string `json:"Nombre"`
	Description string `json:"Descripcion"`
	Contact     string `json:"Contacto"`
}

// SeveralJSONNodes es un conjunto de JSONNodes
type SeveralJSONNodes struct {
	Nodes []JSONNodes `json:"Tiendas"`
}

// SpecificStore es un tipo que permite buscar por medio de 3 parametros
type SpecificStore struct {
	Departament string `json:"Departamento"`
	Name        string `json:"Nombre"`
	Rating      int    `json:"Calificacion"`
}

// VectorItem es un item para llenar el vector linealizado
type VectorItem struct {
	Department string
	Rating     int
	List       *DoublyLinkedList
}

package main

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

package main

import "fmt"

type Robot struct {
	posicion string
	entrega string
	posicionActual string
	encontrados []*Product
	buscados []*Product
	ruta []int
}

func NewRobot(posicion, entrega string) *Robot {
	return &Robot{
		posicion: posicion,
		entrega:  entrega,
		posicionActual: posicion,
	}
}

func (r *Robot) buscarRuta(productos []Product) {
	fmt.Println("$$$ Buscando ruta para los productos asignados")
	for len(r.encontrados) < len(productos)	 {
		caminoInicial := NewCaminoMinimo(*grafo, grafo.obtenerIndice(r.posicion))
		r.siguienteNodo(caminoInicial)
	}
	
}

func (r *Robot) siguienteNodo(camino *CaminoMinimo)  {
	fmt.Println()
	nodosCercanos := camino.NodosMasCercanos()
	for _, cercano := range nodosCercanos {
		nodo := grafo.obtenerNodo(cercano)
		encontrado, producto := r.estaEnProductos(nodo)
		if encontrado{
			r.ruta = append(r.ruta, nodo.id)
			r.posicionActual = nodo.contenido.nombre
			r.encontrados = append(r.encontrados, producto)

		}
	}
}

func (r *Robot) estaEnProductos(nodo *nodoGrafo) (bool, *Product) {
	for _, buscado := range r.buscados {
		if buscado.Almacenamiento == nodo.contenido.nombre && !r.yaEncontrado(buscado) {
			return true, buscado
		}
	}
	return false, nil
}

func (r *Robot) yaEncontrado(product *Product) bool {
	for _, encontrado := range r.encontrados {
		if encontrado == product {
			return true
		}
	}
	return false
}

func (r *Robot) imprimirRuta() {
	for _, encontrado := range r.encontrados {
		fmt.Print(encontrado, " -> ")
	}
	fmt.Println()
}
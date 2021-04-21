package main

import (
	"fmt"
	"strconv"
)

type Robot struct {
	posicion string
	entrega string
	ruta []int
}

func NewRobot(posicion, entrega string) *Robot {
	return &Robot{
		posicion: posicion,
		entrega:  entrega,
	}
}

func (r *Robot) buscarRuta(productos []Product) {
	fmt.Println("$$$ Buscando ruta para los productos asignados")

	inicio := grafo.obtenerIndice(r.posicion)
	r.ruta = append(r.ruta, inicio)

	var paradas []string
	for _, producto := range productos {
		paradas = append(paradas, producto.Almacenamiento)
	}
	despacho := grafo.obtenerIndice("Despacho")
	inico := grafo.obtenerIndice(r.posicion)
	nodos := grafo.obtenerIndices(paradas)
	rute := grafo.rutaConParadas(inico, nodos)
	rute = append(rute, despacho)
	r.ruta = append(r.ruta, rute... )

	r.ruta = append(r.ruta, inicio)

}


func (r *Robot) imprimirRuta(g *Grafo) {

	for _, encontrado := range r.ruta {
		fmt.Print(g.obtenerNodo(encontrado).contenido.nombre, " -> ")
	}
	fmt.Println()
	r.toDot(g)
}

func (r *Robot) toDot(g *Grafo) {
	var texto string
	nodos := g.obtenerNodos(r.ruta)
	for _, nodo := range nodos {
		texto += "\t" + nodo.contenido.nombre + "[shape=\"circle\"];\n"
	}
	var anterior *nodoGrafo
	for i, n := range nodos {
		if i != 0 {
			//arista := g.obtenerArista(anterior.id, n.id)
			/*if arista != nil {
				//texto +=  anterior.contenido.nombre + " ->" + n.contenido.nombre + "[label= \""+strconv.Itoa(arista.peso)+"\"];\n"

			}*/
			texto +=  anterior.contenido.nombre + " ->" + n.contenido.nombre + "[label= \""+strconv.Itoa(i)+"\"];\n"
		}
		anterior = n
	}
	fmt.Println(texto)
}
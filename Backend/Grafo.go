package main

import (
	"fmt"
	"strconv"
)

type contenidoGrafo struct {
	nombre string
}

type nodoGrafo struct {
	id int
	contenido contenidoGrafo
	adyacentes []*nodoGrafo
}

type arista struct {
	peso int
	anterior *nodoGrafo
	siguiente *nodoGrafo
}

func NewArista (peso int, siguiente, anterior *nodoGrafo) *arista{
	return &arista{
		peso: peso,
		siguiente: siguiente,
		anterior : anterior,
	}
}

type Grafo struct {
	nodos []*nodoGrafo
	aristas []*arista
}

func NewGrafo() *Grafo {
	return &Grafo{}
}

func (g *Grafo) agregarNodo(id int, contenido contenidoGrafo)  {

	if !estaEn(g.nodos, id){
		g.nodos = append(g.nodos, &nodoGrafo{id : id, contenido : contenido})
	}
}

func (g *Grafo) agregarArista(anterior, siguiente int , peso int)  {
	nodoAnterior := g.obtenerNodo(anterior)
	nodoSig := g.obtenerNodo(siguiente)
	if nodoAnterior != nil && nodoSig != nil {
		 if !estaEn(nodoAnterior.adyacentes, nodoSig.id){
			 nodoAnterior.agregarAdyacente(nodoSig)
			 g.aristas = append(g.aristas, NewArista(peso, nodoSig, nodoAnterior))
		 } else {
			 fmt.Errorf("no se agrego la arista")
		 }
	} else {
		fmt.Errorf("no se agrego la arista")
	}
}

func (g *Grafo) obtenerIndice(nombre string) int {
	var indice int

	for _, n := range g.nodos {
		if n.contenido.nombre == nombre {
			indice = n.id
		}
	}

	return indice
}

func (n *nodoGrafo) agregarAdyacente(nuevo *nodoGrafo)  {
	n.adyacentes = append(n.adyacentes, nuevo)
}

func (g *Grafo) obtenerNodo(id int) *nodoGrafo {
	for i, nodo := range g.nodos {
		if nodo.id == id {
			return g.nodos[i]
		}
	}
	return  nil
}

func (g *Grafo) obtenerNodos(id []int) []*nodoGrafo {
	var nodos []*nodoGrafo
	for _,i := range id {
		nodos = append(nodos, g.obtenerNodo(i))
	}
	return nodos
}

func (g *Grafo) obtenerArista(inicio int, fin int) *arista {
	for i, a := range g.aristas {
		if a.anterior.id == inicio && a.siguiente.id == fin{
			return g.aristas[i]
		}
	}
	return nil
}

func estaEn ( lista []*nodoGrafo, id int) bool {
	for _, i := range lista {
		if id == i.id {
			return true
		}
	}
	return false
}

func (g Grafo) Imprimir() {
	for _, i := range g.nodos {
		fmt.Print("Nodo ",i.id, " : ")
		for _, j := range i.adyacentes {
			fmt.Print(j.id)
		}
		fmt.Println()
	}
}

func (g *Grafo) obtenerIndices(lista []string) []int {
	var indices []int
	for _, l := range lista {
		indices = append(indices,  g.obtenerIndice(l))
	}
	return indices
}

func (g *Grafo) peso(inicio, fin int) int {
	arista := g.obtenerArista(inicio, fin)

	if arista != nil {
		return arista.peso
	} else {
		return -1
	}
}

func (g Grafo) toDot() string {
	var texto string

	for _, nodo := range g.nodos {
		texto += "\t" + nodo.contenido.nombre + "[shape=\"circle\"];\n"
	}

	for _, a := range g.aristas {
		texto +=  a.anterior.contenido.nombre + " ->" + a.siguiente.contenido.nombre + "[label= \""+strconv.Itoa(a.peso)+"\"];\n"
	}

	return texto

}

func (grafo *Grafo) graficarGrafo()  {
	texto := "digraph grafo { \n\tnode[shape=\"record\" style=\"filled\" fillcollor=\"#58D27A\"]\n"
	texto += grafo.toDot()
	texto += "\n}"
	escribirDOT(texto, "Grafo")
	ejecutarComand("CaminoRobot")

	fmt.Println(texto)
}

func (grafo *Grafo) generarCamino(inicio, fin int) *CaminoMinimo {
	fmt.Print("$$$ Calculando camino minimo de ", inicio, " a ", fin)
	camino := NewCaminoMinimo(*grafo, inicio)
	camino.ultimo = append(camino.ultimo, inicio)
	fmt.Print("$$$ El camino es : ")
	fmt.Print(inicio)
	camino.generarCamino(*grafo,*grafo.nodos[inicio], *grafo.nodos[fin])
	camino.ultimo = append(camino.ultimo, fin)
	return camino
}

func (g *Grafo) rutaConParadas(inicio int, paradas []int) []int {
	actual := inicio
	var encontrado []int
	var ruta []int
	for i, n := range g.nodos {
		fmt.Println(i, ". ", n.contenido.nombre)
	}
	for len(encontrado) < len(paradas)-1 {

		encontrado = append(encontrado, actual)

		camino := NewCaminoMinimo(*g, actual)
		camino.Dijkstra(*g, *g.obtenerNodo(actual))
		cercanos := camino.NodosMasCercanos()
		siguiente := camino.siguienteNodo(paradas, encontrado, cercanos)
		camino.generarCamino(*grafo, *grafo.obtenerNodo(actual), *grafo.obtenerNodo(siguiente))
		camino.camino = append(camino.camino, siguiente)
		siguientes := camino.camino
		siguientes = append(siguientes[:])
		ruta = append(ruta, siguientes...)
		actual = siguiente

	}
	return ruta
}


/*func main() {
	grafo := NewGrafo()
	contenido1 := &contenidoGrafo{"Despacho"}
	contenido2 := &contenidoGrafo{"Ferreteria"}
	contenido3 := &contenidoGrafo{"Miscelania"}
	contenido4 := &contenidoGrafo{"Hogar"}
	contenido5 := &contenidoGrafo{"Pintura"}

	grafo.agregarNodo(0, *contenido1)
	grafo.agregarNodo(1, *contenido2)
	grafo.agregarNodo(2, *contenido3)
	grafo.agregarNodo(3, *contenido4)
	grafo.agregarNodo(4, *contenido5)

	grafo.agregarArista(0, 1, 50)
	grafo.agregarArista(0, 2, 100)
	grafo.agregarArista(0, 2, 200)
	grafo.agregarArista(1, 2, 25)
	grafo.agregarArista(1, 3, 80)
	grafo.agregarArista(2,3, 50)
	grafo.agregarArista(2,4, 10)
	grafo.agregarArista(0,4, 500)


	grafo.generarCamino(0, 4)
}*/


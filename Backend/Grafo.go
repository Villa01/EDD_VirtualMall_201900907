package main

import (
	"fmt"
	"strconv"
)

type contenidoGrafo struct {
	
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

func (g *Grafo) agregarArista(anterior, siguiente, peso int)  {
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

func (g Grafo) toDot() string {
	var texto string

	for _, nodo := range g.nodos {
		texto += "\tN" + strconv.Itoa(nodo.id) + "[shape=\"circle\"];\n"
	}

	for _, a := range g.aristas {
		texto += "N" + strconv.Itoa(a.anterior.id) + " -> N" + strconv.Itoa(a.siguiente.id) + "[label= \""+strconv.Itoa(a.peso)+"\"];\n"
	}

	return texto

}

/*func (g *Grafo) rutaMinima(inicio *nodoGrafo, fin *nodoGrafo) []*nodoGrafo{
	var visitados []*nodoGrafo
	noVisitados := g.nodos
	actual := inicio


}*/


func (g *Grafo) distanciaEntre(inicio *nodoGrafo, fin *nodoGrafo, valorInicio int) int{
	for _, a := range g.aristas {
		if a.anterior == inicio &&  a.siguiente== fin{
			return a.peso
		}
	}
	return 1000000000000000000000000000
}



func main() {
	prueba := NewGrafo()

	for i := 0; i < 8; i++ {
		var contenido contenidoGrafo
		prueba.agregarNodo(i, contenido)
	}

	prueba.agregarArista(0, 1, 100)

	prueba.agregarArista(0, 2, 200)

	prueba.agregarArista(0, 3, 150)

	prueba.agregarArista(0, 4, 25)

	prueba.agregarArista(0, 1, 10)
	prueba.agregarArista(1, 2, 5)
	prueba.agregarArista(2, 3, 40)
	prueba.agregarArista(3, 4, 80)
	prueba.agregarArista(4, 5, 95)
	prueba.agregarArista(5, 6, 99)
	prueba.agregarArista(6, 7, 101)
	prueba.agregarArista(7, 0, 3004)

	fmt.Println(prueba.toDot())
	prueba.Imprimir()
}
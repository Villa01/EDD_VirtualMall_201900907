package main

import (
	"fmt"
	"strconv"
)

type contenidoGrafo struct {
	nombre string
}

type nodoGrafo struct {
	id string
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

func (g *Grafo) agregarNodo(id string, contenido contenidoGrafo)  {

	if !estaEn(g.nodos, id){
		g.nodos = append(g.nodos, &nodoGrafo{id : id, contenido : contenido})
	}
}

func (g *Grafo) agregarArista(anterior, siguiente string , peso int)  {
	nodoAnterior := g.obtenerNodo(anterior)
	nodoSig := g.obtenerNodo(siguiente)
	fmt.Println("anterior ", nodoAnterior)
	fmt.Println("siguiente ", nodoSig)
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

func (g *Grafo) obtenerNodo(id string) *nodoGrafo {
	for i, nodo := range g.nodos {
		if nodo.id == id {
			return g.nodos[i]
		}
	}
	return  nil
}

func estaEn ( lista []*nodoGrafo, id string) bool {
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
		texto += "\tN" + nodo.id + "[shape=\"circle\"];\n"
	}

	for _, a := range g.aristas {
		texto +=  a.anterior.id + " ->" + a.siguiente.id + "[label= \""+strconv.Itoa(a.peso)+"\"];\n"
	}

	return texto

}

/*func (g *Grafo) rutaMinima(inicio *nodoGrafo, fin *nodoGrafo) []*nodoGrafo{
	var visitados []*nodoGrafo
	noVisitados := g.nodos
	actual := inicio


}*/


//func (g *Grafo) distanciaEntre(inicio *nodoGrafo, fin *nodoGrafo, valorInicio int) int{
//	for _, a := range g.aristas {
//		if a.anterior == inicio &&  a.siguiente== fin{
//			return a.peso
//		}
//	}
//	return 1000000000000000000000000000
//}




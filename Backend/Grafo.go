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
}

func (grafo *Grafo) generarCamino(inicio, fin int) *CaminoMinimo {
	fmt.Print("$$$ Calculando camino minimo de ", inicio, " a ", fin)
	camino := NewCaminoMinimo(*grafo, inicio)
	camino.ultimo = append(camino.ultimo, inicio)
	fmt.Print("$$$ El camino es : ")
	fmt.Print(inicio)
	camino.Dijkstra(*grafo,*grafo.nodos[inicio], *grafo.nodos[fin])
	camino.ultimo = append(camino.ultimo, fin)
	fmt.Print("-> ", fin )
	return camino
}


type CaminoMinimo struct {
	ultimo []int
	D []int
	V []bool
	n, s int

}

func NewCaminoMinimo( g Grafo, origen int) *CaminoMinimo {

	infinito = 1000000000000
	var ultimo []int
	var visto []bool
	var D []int

	for i:= 0; i<  len(g.nodos); i++ {
		ultimo = append(ultimo, 0)
	}

	for i:= 0; i<  len(g.nodos); i++ {
		D = append(D, 1000000000000000000)
	}

	for i:= 0; i<  len(g.nodos); i++ {
		visto = append(visto, false)
	}
	D[origen] = 0

	return &CaminoMinimo{
		ultimo: ultimo,
		D:      D,
		V: 		visto,
		s:      origen,
		n:      len(g.nodos),
	}
}

var infinito int
func (c *CaminoMinimo) Dijkstra(g Grafo,  s nodoGrafo, f nodoGrafo) {
	for i, w := range g.nodos {
		if g.obtenerArista(s.id, w.id) == nil {
			c.D[w.id] = 1000000000000000000 // Valor de infinito
		} else {
			peso := g.peso(s.id, w.id)
			if peso != -1 {
				c.D[w.id] = peso
			}
		}
		c.ultimo[i] = s.id
	}
	c.D[s.id] = 0
	c.V[s.id] = true
	contador := 0
	for !c.todosVistos() {

		vertice := c.minimo(contador)
		if vertice == -1 {
			break
		}
		if vertice != f.id {
			c.ultimo = append(c.ultimo, vertice)
			fmt.Print("-> ", vertice)

			c.V[vertice] = true
			for _, nodo := range g.obtenerNodo(vertice).adyacentes {
				peso := g.peso(vertice, nodo.id)
				if peso != -1 && c.D[ nodo.id] > c.D[vertice] + peso{
					c.D[ nodo.id] = c.D[vertice] + peso
				}
			}
			contador ++
		} else {
			return
		}
	}
}


func (c *CaminoMinimo) todosVistos() bool {
	for _, v := range c.V {
		if v == false {
			return false
		}
	}
	return true
}

func (c *CaminoMinimo) minimo(anterior int) int {
	minimo := infinito
	retorno := -1
	for i, d := range c.D {
		if d < minimo && !c.V[i] {
			minimo = d
			retorno = i
		}
	}
	if retorno == -1 {
		c.V[anterior+1] = true
	}
	return retorno
}

// Retorna los nodos mas cercanos
func (c *CaminoMinimo) NodosMasCercanos() []int{
	var retorno []int
	distancias := c.D

	for i, _ := range distancias {
		retorno = append(retorno, i)
	}
	num := len(c.D)
	for z := 1; z < num; z++ {
		for v := 0; v < (num - z); v++ {
			if distancias[v] > distancias[v+1]{
				aux := retorno[v]
				retorno[v] = retorno[v + 1]
				retorno[v + 1] = aux
			}
		}
	}

	return retorno
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
}
*/

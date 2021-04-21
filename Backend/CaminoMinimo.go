package main

import "fmt"

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
func (c *CaminoMinimo) generarCamino(g Grafo,  s nodoGrafo, f nodoGrafo) {
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

func (c *CaminoMinimo) Dijkstra(g Grafo,  s nodoGrafo) {
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
		c.ultimo = append(c.ultimo, vertice)

		c.V[vertice] = true
		for _, nodo := range g.obtenerNodo(vertice).adyacentes {
			peso := g.peso(vertice, nodo.id)
			if peso != -1 && c.D[ nodo.id] > c.D[vertice] + peso{
				c.D[ nodo.id] = c.D[vertice] + peso
			}
		}
		contador ++
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
	var aux2 int
	var aux int
	for z := 0; z < len(distancias); z++ {
		for v := z+1; v < len(distancias); v++ {
			if distancias[z] > distancias[v]{
				aux = distancias[z]
				distancias[z] = distancias[v]
				distancias[v] = aux

				aux2 = retorno[z]
				retorno[z] = retorno[v]
				retorno[v] = aux2
			}
		}
	}
	return retorno
}

func (c *CaminoMinimo) siguienteNodo(paradas, encontrados, cercanos []int ) int {
	var siguiente int
	for _, cercano := range cercanos {
		// Verificar que esté dentro de las paradas
		if c.estaEn(cercano, paradas) {
			// Verificar que no ha sido encontrado aun
			if !c.estaEn(cercano, encontrados) && cercano != 0{
				fmt.Println("Encontre uno cercano ", cercano)
				siguiente = cercano
				break
			}
		}
	}
	// Si no encontró una parada se va a la mas cercana
	if siguiente == 0 {
		siguiente = paradas[1]
		fmt.Println("Encontre la primera mas cercana ", siguiente)
	}
	return siguiente
}

func (c *CaminoMinimo) estaEn(elemento int, lista []int) bool {
	for i := 0; i < len(lista); i++ {
		if elemento == lista[i]{
			return true
		}
	}

	return false
}

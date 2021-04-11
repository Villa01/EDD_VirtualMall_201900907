package main

import (
	"fmt"
	"log"
)

// Implementación del nodo del arbol B

type Clave struct {
	indice int
	cuenta *Cuenta
}

// NewClave inicializa una nueva clave dentro de las paginas
func NewClave(indice int, cuenta *Cuenta) *Clave {
	return &Clave{indice , cuenta }
}

// Pagina es un nodo del árbol B
type Pagina struct {
	claves []Clave
	ramas []*Pagina
	cuenta int
	maximo int
}

// NewPagina inicializa un nuevo nodo del Arbol B
func NewPagina( orden int ) *Pagina {
	var claves []Clave
	var ramas []*Pagina

	for i:= 0; i<orden; i++ {
		ramas[i] = nil
	}
	return &Pagina{claves,  ramas, 0, orden }
}

// nodoLleno verifica que si un nodo llego a su capacidad máxima
func (p Pagina) nodoLleno() bool {
	return p.cuenta == p.maximo -1
}

// nodoSemiVacio verifica si el nodo tiene su cantidad minima de claves
func (p Pagina) nodoSemiVacio() bool {
	return p.cuenta < p.maximo/2
}

// dameClave obtiene la clave en cierta posicion
func (p *Pagina) dameClave(posicion int) *Clave {
	return &p.claves[posicion]
}

// cambiarClave modifica la clave en cierta posición
func (p *Pagina) cambiarClave(posicion int, nueva Clave)  {
	p.claves[posicion] = nueva
}

// dameRama obtiene la rama en cierta posicion
func (p *Pagina) dameRama(posicion int) *Pagina {
	return p.ramas[posicion]
}

// cambiarRama modifica la rama en cierta posicion
func (p *Pagina) cambiarRama(posicion int, nueva *Pagina) {
	p.ramas[posicion] = nueva
}

// dameCuenta retorna la cuenta de claves en el nodo
func (p *Pagina) dameCuenta() int {
	return p.cuenta
}

func (p *Pagina) cambiarCuenta(i int) {
	p.cuenta = i
}

// Implementación de Arbol B

// ArbolB es un tipo de arbol de ordenamiento por indices
type ArbolB struct {
	orden int
	raiz *Pagina
}

// NewArbolB crea un nuevo ArbolB con cierto orden
func NewArbolB( orden int ) *ArbolB {
	return &ArbolB{orden, nil}
}

// arbolBVacio verifica si el arbol no tiene nodos
func (a ArbolB) arbolBVacio() bool {
	return a.raiz == nil
}

// dameRaiz retorna la raiz del ArbolB
func (a *ArbolB) dameRaiz() *Pagina {
	return a.raiz
}

// cambiarRaiz cambia la raiz del ArbolB
func (a *ArbolB) cambiarRaiz(nueva *Pagina)  {
	a.raiz = nueva
}

// dameOrden retorna el orden del ArbolB
func (a *ArbolB) dameOrden() int {
	return a.orden
}

// cambiarOrden cambia el orden del ArbolB
func (a *ArbolB) cambiarOrden(nuevo int) {
	a.orden = nuevo
}

// buscarClave busca una clave en cierta pagina, devuelve true si la encontro y la posicion en la pagina
func (a *ArbolB) buscarClave(actual *Pagina, clave Clave) (bool, int) {
	var index int
	var encontrado bool

	// Busca en la primera posicion porque el orden es descendente
	if clave.indice < actual.dameClave(1).indice { // Las claves se almacenan desde la posicion 1
		encontrado = false
		index = 0
	} else {
		index = actual.dameCuenta()
		for clave.indice < actual.dameClave(index).indice && index >1 {
			index--
			encontrado = clave.indice == actual.dameClave(index).indice
		}
	}
	return encontrado, index

}

// Buscar funcion publica para encontrar una clave en el arbolB
func (a *ArbolB) Buscar(clave Clave) (*Pagina, int) {
	return a.buscar(a.dameRaiz(), clave )
}

// buscar funcion recursiva para encontrar una clave en el arbolB
func (a *ArbolB) buscar(actual *Pagina, clave Clave) (*Pagina, int) {
	if actual == nil {
		 return nil, 0
	} else {
		var encontrado, n = a.buscarClave(actual, clave)

		if encontrado {
			return actual, n
		} else {
			return a.buscar(actual.dameRama(n), clave)
		}
	}
}

func (a *ArbolB) Insertar(nueva Clave) {
	a.cambiarRaiz(a.insertar(a.raiz, nueva))
}

func (a *ArbolB) insertar(raiz *Pagina, clave Clave) *Pagina {
	var sube bool
	var mediana Clave
	var nd *Pagina

	sube, mediana, nd = a.empujar(raiz, clave)

	if sube { // Crea un nuevo nodo, con referencia a los otros dos nodos
		p := NewPagina(a.dameOrden())
		p.cambiarCuenta(1)
		p.cambiarClave(1, mediana)
		p.cambiarRama(0, raiz)
		p.cambiarRama(1, nd)
		a.raiz = p
	}
	return raiz
}

func (a *ArbolB) empujar(actual *Pagina, clave Clave) (bool, Clave, *Pagina){

	var k int
	sube := false
	var mediana Clave
	var nuevo *Pagina


	if actual == nil {
		sube = true
		mediana = clave
		nuevo = nil
	}else {
		var esta bool
		esta,k = a.buscarClave(actual, clave)

		if (esta) {
			log.Fatal("La clave esta dublicada")
		}
		sube, mediana, nuevo = a.empujar(actual.dameRama(k), clave)

		if sube {
			if actual.nodoLleno() {
				mediana, nuevo = a.dividirNodo(actual, &mediana, nuevo, k)
			}else {
				sube = false
				a.meterPagina(actual, mediana, nuevo, k)
			}
		}
	}
	return sube, mediana, nuevo
}

func (a *ArbolB) meterPagina(actual *Pagina, clave Clave, ramaDr * Pagina, k int)  {
	for i := actual.dameCuenta(); i>= k +1; i-- {
		actual.cambiarClave(i+1, *actual.dameClave(i))
		actual.cambiarRama(i + 1 , actual.dameRama(i))
	}
	actual.cambiarClave(k+1, clave)
	actual.cambiarRama(k+1, ramaDr)
	actual.cambiarCuenta(actual.dameCuenta()+1)
}

func (a *ArbolB) dividirNodo(actual *Pagina, mediana *Clave, nuevo * Pagina, pos int) (Clave, *Pagina) {
	var i, posMdna, k int
	var nuevaPagina *Pagina
	orden := a.dameOrden()
	k = pos
	if k <= orden/2 {
		k = orden/2
	} else {
		k = orden + 1
	}
	nuevaPagina = NewPagina(orden)
	for i = posMdna +1; i < orden; i++ {
		nuevaPagina.cambiarClave(i-posMdna, *actual.dameClave(i))
		nuevaPagina.cambiarRama(i-posMdna,actual.dameRama(i))

	}

	nuevaPagina.cambiarCuenta((orden-1)-posMdna)
	actual.cambiarCuenta(posMdna)

	if k <= orden/2 {
		a.meterPagina(actual, *mediana, nuevo, pos)
	} else {
		pos = k - posMdna
		a.meterPagina(nuevaPagina, *mediana, nuevo, pos)
	}
	mediana = actual.dameClave(actual.dameCuenta())
	nuevaPagina.cambiarRama(0, actual.dameRama(actual.dameCuenta()))
	actual.cambiarCuenta(actual.dameCuenta()-1)
	return *mediana, nuevo
}

func (a *ArbolB) Escribir()  {
	a.escribir(a.raiz, 1)
}

func (a ArbolB) escribir(r *Pagina, h int)  {
	var i int
	if r != nil {
		a.escribir(r.dameRama(0), h + 5)
		for i := 1; i<= r.dameCuenta()/2; i++ {
			a.escribir(r.dameRama(i), h+5)
			fmt.Println()
		}
		for i=1; i<= r.dameCuenta(); i++ {
			for  j := 0; j <= h; j++{
				fmt.Println(" ")
			}
			fmt.Println(r.dameClave(i))
		}
		for i := r.dameCuenta()/2 +1; i <= r.dameCuenta(); i++ {
			a.escribir(r.dameRama(i), h+ 5)

		}
		fmt.Println()
	}

}

func main() {
	arbol := NewArbolB(5)
	cuenta1 := NewCuenta(1)
	cuenta2 := NewCuenta(2)
	cuenta3 := NewCuenta(3)
	cuenta4 := NewCuenta(4)
	cuenta5 := NewCuenta(5)
	cuenta6 := NewCuenta(6)

	clave1 := NewClave(cuenta1.DPI, cuenta1)
	clave2 := NewClave(cuenta2.DPI, cuenta2)
	clave3 := NewClave(cuenta3.DPI, cuenta3)
	clave4 := NewClave(cuenta4.DPI, cuenta4)
	clave5 := NewClave(cuenta5.DPI, cuenta5)
	clave6 := NewClave(cuenta6.DPI, cuenta6)

	arbol.Insertar(*clave1)
	arbol.Insertar(*clave2)
	arbol.Insertar(*clave3)
	arbol.Insertar(*clave4)
	arbol.Insertar(*clave5)
	arbol.Insertar(*clave6)

	arbol.Escribir()

}

// TODO: Eliminar()


// TODO: Graficar()

package main

// Implementación del nodo del arbol B

type Clave struct {
	indice int
	cuenta *Cuenta
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

// TODO: Insertar()

// TODO: Eliminar()

// TODO: Buscar()

// TODO: Graficar()

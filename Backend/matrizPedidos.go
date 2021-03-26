package main

import "fmt"



type nodoM struct {
	x,y int
	cola *ColaPedidos
	izquierda, derecha, arriba, abajo *nodoM

	header int
	siguiente, anterior *nodoM
}

type lista struct {
	first, last *nodoM
}

type matriz struct {
	lst_h, lst_v *lista
}


func nodoMatriz(x int, y int, cola *ColaPedidos) *nodoM {
	return &nodoM{x,y,cola, nil,nil,nil,nil,0,nil,nil}
}

func nodoLista(header int) *nodoM{
	return &nodoM{0,0,nil,nil,nil,nil,nil,header,nil, nil}
}

func newLista() *lista{
	return &lista{nil,nil}
}

//Se cambio a primer letra mayuscula para poder acceder
func NewMatriz() *matriz{
	return &matriz{newLista(),newLista()}
}

func (n *nodoM) headerX() int { return n.x }

func (n *nodoM) headerY() int { return n.y }
//+ n.cola. +
func (n *nodoM) toString() string{ return "Nombre: \n"}

func (l *lista ) ordenar(nuevo *nodoM)  {
	aux := l.first
	for(aux != nil){
		if nuevo.header > aux.header{
			aux = aux.siguiente
		}else{
			if aux == l.first{
				nuevo.siguiente = aux
				aux.anterior = nuevo
				l.first = nuevo
			}else{
				nuevo.anterior = aux.anterior
				aux.anterior.siguiente = nuevo
				nuevo.siguiente = aux
				aux.anterior = nuevo
			}
			return
		}
	}
	l.last.siguiente = nuevo
	nuevo.anterior = l.last
	l.last = nuevo
}

func (l *lista) insert(header int) {
	nuevo := nodoLista(header)
	if l.first == nil{
		l.first = nuevo
		l.last = nuevo
	}else{
		l.ordenar(nuevo)
	}
}

func (l *lista) search(header int) *nodoM{
	temp := l.first
	for temp != nil{
		if temp.header == header{
			return temp
		}
		temp = temp.siguiente
	}
	return nil
}

func (l *lista) print() {
	temp := l.first
	for temp != nil{
		fmt.Println("Cabecera:", temp.header)
		temp = temp.siguiente
	}
}

func (m *matriz) Insert(meses *ColaPedidos, x int, y int){
	h := m.lst_h.search(x)
	v := m.lst_v.search(y)

	if h==nil && v==nil{
		m.noExisten(meses, x,y)
	}else if h==nil && v!=nil{
		m.existeVertical(meses, x, y)
	}else if h!=nil && v==nil{
		m.existeHorizontal(meses, x, y)
	}else{
		m.existen(meses,x,y)
	}
}

func (m *matriz)noExisten(meses *ColaPedidos, x int, y int) {
	m.lst_h.insert(x)
	m.lst_v.insert(y)

	h := m.lst_h.search(x)
	v := m.lst_v.search(y)

	nuevo := nodoMatriz(x,y,meses)

	h.abajo = nuevo
	nuevo.arriba = h

	v.derecha = nuevo
	nuevo.izquierda = v
}

func (m *matriz) existeVertical(meses *ColaPedidos, x int, y int) {
	m.lst_h.insert(x)
	h := m.lst_h.search(x)
	v := m.lst_v.search(y)

	nuevo := nodoMatriz(x,y, meses)
	agregado := false

	aux := v.derecha

	var cabecera int

	for aux!=nil{
		cabecera = aux.headerX()
		if cabecera > x{
			aux = aux.derecha
		} else {
			nuevo.derecha = aux
			nuevo.izquierda = aux.izquierda
			aux.izquierda.derecha = nuevo
			aux.izquierda = nuevo
			agregado = true
			break
		}

	}
	if !agregado {
		aux = v.derecha
		for aux.derecha != nil {
			aux = aux.derecha
		}
		nuevo.izquierda = aux
		aux.derecha = nuevo
	}

	nuevo.arriba = h
	h.abajo = nuevo

}

func (m *matriz) existeHorizontal(meses *ColaPedidos, x int, y int) {
	m.lst_v.insert(y)
	h := m.lst_h.search(x)
	v := m.lst_v.search(y)

	nuevo := nodoMatriz(x,y, meses)
	agregado := false

	aux := h.abajo

	var cabecera int

	for aux!=nil{
		cabecera = aux.headerY()
		if cabecera > y{
			aux = aux.abajo
		} else {
			nuevo.abajo = aux
			nuevo.arriba = aux.arriba
			aux.arriba.abajo = nuevo
			aux.arriba = nuevo
			agregado = true
			break
		}

	}
	if !agregado {
		aux = v.abajo
		for aux.abajo != nil {
			aux = aux.abajo
		}
		nuevo.arriba = aux
		aux.abajo = nuevo
	}

	nuevo.izquierda = v
	v.derecha = nuevo

}

func (m *matriz) existen(meses *ColaPedidos, x int, y int) {
	h := m.lst_h.search(x)
	v := m.lst_v.search(y)

	nuevo := nodoMatriz(x,y, meses)
	agregado := false

	aux := v.derecha

	var cabecera int

	for aux!=nil{
		cabecera = aux.headerX()
		if cabecera > x{
			aux = aux.derecha
		} else {
			nuevo.derecha = aux
			nuevo.izquierda = aux.izquierda
			aux.izquierda.derecha = nuevo
			aux.izquierda = nuevo
			agregado = true
			break
		}

	}
	if !agregado {
		aux = v.derecha
		for aux.derecha != nil {
			aux = aux.derecha
		}
		nuevo.izquierda = aux
		aux.derecha = nuevo
	}

	agregado = false

	aux = h.abajo

	for aux!=nil{
		cabecera = aux.headerY()
		if cabecera > y{
			aux = aux.abajo
		} else {
			nuevo.abajo = aux
			nuevo.arriba = aux.arriba
			aux.arriba.abajo = nuevo
			aux.arriba = nuevo
			agregado = true
			break
		}

	}
	if !agregado {
		aux = v.abajo
		for aux.abajo != nil {
			aux = aux.abajo
		}
		nuevo.arriba = aux
		aux.abajo = nuevo
	}
}


func (m matriz) print_vertical(){
	cabecera_y := m.lst_v.first
	for cabecera_y != nil  {
		aux := cabecera_y.derecha
		for aux != nil {
			aux.printLista()
			aux = aux.derecha
		}

		cabecera_y = cabecera_y.abajo
	}

}

func (m *matriz) BuscarNodo(x int, y int) *nodoM{
	var nodo *nodoM

	encabezado := m.lst_h.first

	for encabezado != nil{
		if encabezado.x == x{
			aux := encabezado.abajo
			for aux != nil {
				if aux.y == y{
					nodo = aux
				}
			}
		}
		encabezado = encabezado.siguiente
	}
	return nodo
}

func (m matriz) print_horizontal(){
	cabecera := m.lst_v.first

	for cabecera != nil {
		aux := cabecera.abajo
		for aux != nil {
			aux.printLista()
			aux = aux.abajo
		}
		cabecera = cabecera.siguiente

	}
}

func (n nodoM) printLista() {
	fmt.Println("x: ", n.x , " y: ", n.y)
}


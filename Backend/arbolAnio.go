package main

import (
	"fmt"
)

type nodoPedido struct {
	meses *listaMeses
	indice   int
	altura   int
	izq, der *nodoPedido
}

func newNodoPedido(indice int, meses *listaMeses) *nodoPedido {
	return &nodoPedido{meses,indice, 0, nil, nil}
}

type AVLPedido struct {
	raiz *nodoPedido
}

func NewAVLPedido() *AVLPedido {
	return &AVLPedido{nil}
}

func alturaPedido(temp *nodoPedido) int {
	if temp != nil {
		return temp.altura
	}
	return -1
}

func rotacionIzquierdaP(temp **nodoPedido) {
	aux := (*temp).izq
	(*temp).izq = aux.der
	aux.der = *temp
	(*temp).altura = max(alturaPedido((*temp).der), alturaPedido((*temp).izq)) + 1
	aux.altura = max(alturaPedido((*temp).izq), (*temp).altura) + 1
	*temp = aux
}

func rotacionDerechaP(temp **nodoPedido) {
	aux := (*temp).der
	(*temp).der = aux.izq
	aux.izq = *temp
	(*temp).altura = max(alturaPedido((*temp).der), alturaPedido((*temp).izq)) + 1
	aux.altura = max(alturaPedido((*temp).der), (*temp).altura) + 1
	*temp = aux
}

func rotacionDobleIzquierdaP(temp **nodoPedido) {
	rotacionDerechaP(&(*temp).izq)
	rotacionIzquierdaP(temp)
}

func rotacionDobleDerechaP(temp **nodoPedido) {
	rotacionIzquierdaP(&(*temp).der)
	rotacionDerechaP(temp)
}

func insertP(meses *listaMeses, indice int, root **nodoPedido) {
	if (*root) == nil {
		*root = newNodoPedido(indice, meses )
		return
	}
	if indice < (*root).indice {
		insertP(meses,indice, &(*root).izq)
		if (alturaPedido((*root).izq) - alturaPedido((*root).der)) == -2 {
			if indice < (*root).izq.indice{
				rotacionIzquierdaP(root)
			}else{
				rotacionDobleIzquierdaP(root)
			}
		}
	}else if indice > (*root).indice{
		insertP(meses,indice, &(*root).der)
		if (alturaPedido((*root).der) - alturaPedido((*root).izq)) == 2{
			if indice > (*root).der.indice {
				rotacionDerechaP(root)
			}else{
				rotacionDobleDerechaP(root)
			}
		}
	}else{
		fmt.Println("Ya se inserto el indice")
	}

	(*root).altura = max(alturaPedido((*root).izq), alturaPedido((*root).der))+1
}

func (avl *AVLPedido) InsertarP(meses *listaMeses, indice int) {
	insertP(meses, indice, &avl.raiz)
}

func (avl *AVLPedido) Print(){
	inOrdenP(avl.raiz)
}

func inOrdenP(temp *nodoPedido)  {
	if temp != nil {
		inOrdenP(temp.izq)
		fmt.Println("Nombre: ")
		inOrdenP(temp.der)
	}
}

func (avl *AVLPedido) BuscarNodo (codigo int) * nodoPedido{
	return buscarNodoP(avl.raiz, codigo)
}

func buscarNodoP (temp *nodoPedido, codigo int ) *nodoPedido{
	if temp != nil {
		if temp.indice == codigo {
			return temp
		}
		izq := buscarNodoP(temp.izq, codigo)
		der := buscarNodoP(temp.der, codigo)

		if izq != nil {
			return izq
		} else if der != nil {
			return der
		}
	}
	return nil
}
/*
func (avl *AVL) ObtenerProductos () []Product{
	if avl.raiz == nil {
		return nil
	}
	return obtenerProductos(avl.raiz)
}

func obtenerProductos (temp *nodo) []Product{
	var productos []Product
	if temp != nil {
		productos = append(productos, temp.producto)

		izq := obtenerProductos(temp.izq)
		der := obtenerProductos(temp.der)

		if izq != nil {
			productos = append(productos,izq...)
		}
		if der != nil {
			productos = append(productos, der...)
		}
	}
	return productos
}*/
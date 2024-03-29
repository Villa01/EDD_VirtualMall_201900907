package main

import "fmt"

type nodo struct {
	producto Product
	indice   int
	altura   int
	izq, der *nodo
}

func newNodo(indice int, producto Product) *nodo {
	return &nodo{producto,indice, 0, nil, nil}
}

type AVL struct {
	raiz *nodo
}

func NewAVL() *AVL {
	return &AVL{nil}
}


func altura(temp *nodo) int {
	if temp != nil {
		return temp.altura
	}
	return -1
}

func rotacionIzquierda(temp **nodo) {
	aux := (*temp).izq
	(*temp).izq = aux.der
	aux.der = *temp
	(*temp).altura = max(altura((*temp).der), altura((*temp).izq)) + 1
	aux.altura = max(altura((*temp).izq), (*temp).altura) + 1
	*temp = aux
}

func rotacionDerecha(temp **nodo) {
	aux := (*temp).der
	(*temp).der = aux.izq
	aux.izq = *temp
	(*temp).altura = max(altura((*temp).der), altura((*temp).izq)) + 1
	aux.altura = max(altura((*temp).der), (*temp).altura) + 1
	*temp = aux
}

func rotacionDobleIzquierda(temp **nodo) {
	rotacionDerecha(&(*temp).izq)
	rotacionIzquierda(temp)
}

func rotacionDobleDerecha(temp **nodo) {
	rotacionIzquierda(&(*temp).der)
	rotacionDerecha(temp)
}

func insert(producto Product, indice int, root **nodo) {
	if (*root) == nil {
		*root = newNodo(indice, producto )
		return
	}
	if indice < (*root).indice {
		insert(producto,indice, &(*root).izq)
		if (altura((*root).izq) - altura((*root).der)) == -2 {
			if indice < (*root).izq.indice{
				rotacionIzquierda(root)
			}else{
				rotacionDobleIzquierda(root)
			}
		}
	}else if indice > (*root).indice{
		insert(producto,indice, &(*root).der)
		if (altura((*root).der) - altura((*root).izq)) == 2{
			if indice > (*root).der.indice {
				rotacionDerecha(root)
			}else{
				rotacionDobleDerecha(root)
			}
		}
	}else{
		fmt.Println("Ya se inserto el indice")
	}

	(*root).altura = max(altura((*root).izq), altura((*root).der))+1
}

func (avl *AVL) Insertar(producto Product, indice int) {
	//insert(producto, indice, &avl.raiz)
}

func (avl *AVL) Print(){
	//inOrden(avl.raiz)
}

func inOrden(temp *nodo)  {
	if temp != nil {
		inOrden(temp.izq)
		fmt.Println("Nombre: ", temp.producto.Nombre)
		inOrden(temp.der)
	}
}



func (avl *AVL) BuscarNodo (codigo int) * nodo{
	return buscarNodo(avl.raiz, codigo)
}


func buscarNodo (temp *nodo, codigo int ) *nodo{
	if temp != nil {
		if temp.producto.Codigo == codigo {
			return temp
		}
		izq := buscarNodo(temp.izq, codigo)
		der := buscarNodo(temp.der, codigo)

		if izq != nil {
			return izq
		} else if der != nil {
			return der
		}
	}
	return nil
}



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
}


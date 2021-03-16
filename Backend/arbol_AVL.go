package main

type nodoAVL struct {
	contenido Product
	altura    int
	izq       *nodoAVL
	der       *nodoAVL
}

func newNodoAVL(contenido Product) *nodoAVL {
	return &nodoAVL{contenido: contenido, altura: 0, izq: nil, der: nil}
}

type arbolAVL struct {
	root *nodoAVL
}

func newArbolAVL() *arbolAVL {
	return &arbolAVL{root: nil}
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func (nodo *nodoAVL) getAltura() int {
	if nodo != nil {
		return nodo.altura
	}
	return -1
}

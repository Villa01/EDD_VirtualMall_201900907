package main

import "fmt"

func main() {
	arbol := NuevoBTree(5)
	admin := Cuenta{1, "admin", "", "", "administrado"}
	usuario1 := Cuenta{2, "usuario1", "", "", "usuario"}
	usuario2 := Cuenta{3, "usuario2", "", "", "usuario"}
	usuario3 := Cuenta{4, "usuario3", "", "", "usuario"}
	usuario4 := Cuenta{5, "usuario3", "", "", "usuario"}

	arbol.Insert(admin)
	arbol.Insert(usuario1)
	arbol.Insert(usuario2)
	arbol.Insert(usuario3)
	arbol.Insert(usuario4)
	arbol.ImprimirArbol()
}

type Cuenta struct {
	DPI    int    `json:"Dpi"`
	Nombre string `json:"Nombre"`
	Email  string `json:"Correo"`
	Contra string `json:"Password"`
	Cuenta string `json:"Usuario"`
}

// ArbolB es una estrucutura de datos
type ArbolB struct {
	raiz *BNodo
	t    int
}

// ArbolB retorna un arbol B nuevo
func NuevoBTree(grado int) *ArbolB {
	return &ArbolB{nil, grado}
}

// ImprimirLlaves imprime todas las llaves de los nodos
func (B *ArbolB) ImprimirLlaves() {
	if B.raiz != nil {
		B.raiz.imprimirLlaves()
	}
	fmt.Println()
}

// dameRaiz retorna la raiz del arbol
func (B *ArbolB) dameRaiz() *BNodo {
	return B.raiz
}

//Search busca una cierta cuenta en el arbol
func (B *ArbolB) Search(cuenta Cuenta) *BNodo {
	if B.raiz == nil {
		return nil
	} else {
		return B.raiz.search(cuenta)
	}
}

//search metodo recursivo para buscar en el arbol
func (B *ArbolB) search(cuenta Cuenta) *Cuenta {

	aux := B.Search(cuenta)

	if aux != nil {
		for i := 0; i < aux.numero; i++ {
			if aux.llaves[i].DPI == cuenta.DPI {
				return aux.llaves[i]
			}
		}
	}
	return nil
}

//Insert inserta un nuevo nodo en el arbol
func (B *ArbolB) Insert(cuenta Cuenta) {

	if B.raiz == nil {

		B.raiz = NuevoBNodo(B.t, true)
		B.raiz.llaves[0] = &cuenta
		B.raiz.numero = 1

	} else {

		if B.raiz.numero == 2*B.t-1 {
			aux := NuevoBNodo(B.t, false)
			aux.hijos[0] = B.raiz
			aux.separar(0, B.raiz)

			i := 0
			if aux.llaves[0].DPI < cuenta.DPI {
				i++
			}
			aux.hijos[i].insert(cuenta)
			B.raiz = aux

		} else {
			B.raiz.insert(cuenta)
		}
	}
}

// ImprimirArbol imprime el arbol
func (B *ArbolB) ImprimirArbol() {
	if B.raiz != nil {
		imprimirArbol(*B.raiz)
	}
}

// imprimirArbol metodo recursivo para imprimir el arbol
func imprimirArbol(actual BNodo) {
	if actual.hoja {
		for i := 0; i < actual.numero; i++ {
			fmt.Println("DPI", actual.llaves[i].DPI)
		}

	} else {
		for i := 0; i < actual.numero; i++ {
			if !actual.hoja {
				imprimirArbol(*actual.hijos[i])
			}
		}
	}
}

// Eliminar elimina cierta cuenta
func (B *ArbolB) Eliminar(data Cuenta) {
	if B.raiz != nil {
		B.raiz.eliminar(data)

		if B.raiz.numero == 0 {

			if B.raiz.hoja {
				B.raiz = nil

			} else {
				B.raiz = B.raiz.hijos[0]
			}
		}
	}
}

// Vacio retorna true si no hay nodos en el arbol
func (b *ArbolB) Vacio() bool {
	return b.raiz == nil
}

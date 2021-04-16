package main

import (
	"fmt"
	"strconv"
)

// BNodo es un nodo para el ArbolB
type BNodo struct {
	llaves []*Cuenta `json:"llaves"`
	numero int `json:"numero"`
	hijos  []*BNodo `json:"hijos"`
	grado  int
	hoja   bool
}

type Cuenta struct{
	DPI int `json:"Dpi"`
	Nombre string `json:"Nombre"`
	Email string `json:"Correo"`
	Contra string `json:"Password"`
	Cuenta string `json:"Usuario"`
}

func (c Cuenta) toDOT() string {
	return strconv.Itoa(c.DPI) + "\\n" + c.Nombre + "\\n" + c.Email + "\\n" + c.Contra + "\\n" + c.Cuenta
}

//NuevoBNodo crea un nuevo BNodo y retorna un puntero
func NuevoBNodo(grado int, hoja bool) *BNodo {

	var nuevas []*Cuenta

	for i := 0; i< 2*grado-1;i++{
		nuevas = append(nuevas, nil)
	}

	var nuevos []*BNodo

	for i := 0; i< 2*grado ;i++{
		nuevos = append(nuevos, nil)
	}

	return &BNodo{ nuevas, 0, nuevos, grado, hoja,}
}

// search busca una cuenta dentro de un nodo
func (B *BNodo) search( cuenta Cuenta) *BNodo {

	i := 0
	for i < B.numero && cuenta.DPI > B.llaves[i].DPI {
		i++
	}

	if B.llaves[i].DPI ==  cuenta.DPI  {
		return B
	}

	if B.hoja == true {
		return nil
	}

	return B.hijos[i].search(cuenta)

}

// insert inserta una cuenta a un nodo y crea nuevos nodos si es necesario
func (B *BNodo) insert(cuenta Cuenta) {

	i := B.numero - 1
	if B.hoja == true {

		for i >= 0 && cuenta.DPI < B.llaves[i].DPI {
			B.llaves[i + 1] = B.llaves[i]
			i--
		}

		B.llaves[i + 1] = &cuenta
		B.numero = B.numero + 1
	} else {

		for i >= 0 && B.llaves[i].DPI > cuenta.DPI {
			i--
		}

		if B.hijos[i + 1].numero == (2 * B.grado - 1) {
			B.separar(i + 1, B.hijos[i + 1])
			if B.llaves[i + 1].DPI < cuenta.DPI {
				i++
			}
		}
		B.hijos[i + 1].insert(cuenta)
	}
}

// separar separa un nodo en otros 2 nodos
func (B *BNodo) separar( indice int , node *BNodo) {

	aux := NuevoBNodo(node.grado, node.hoja)

	aux.numero = B.grado - 1

	for  j := 0; j < B.grado- 1; j++{
		aux.llaves[j] = node.llaves[j + B.grado]
	}

	if node.hoja == false{
		for j := 0; j < B.grado; j++ {
			aux.hijos[j] = node.hijos[j + B.grado]
		}
	}

	node.numero = B.grado - 1

	for j:= B.numero; j >= (indice + 1); j-- {
		B.hijos[j + 1] = B.hijos[j]
	}

	B.hijos[indice+ 1] = aux

	for j := B.numero - 1; j >= indice; j-- {
		B.llaves[j + 1] = B.llaves[j]
	}

	B.llaves[indice] = node.llaves[B.grado- 1]

	B.numero = B.numero + 1

}

// buscarllave busca la llave en un nodo
func (B *BNodo) buscarllave( cuenta Cuenta) int {
	index := 0
	for index < B.numero && B.llaves[index].DPI < cuenta.DPI {
		index++
	}
	return index
}

// eliminar elimina una cuenta de un nodo
func (B *BNodo) eliminar(cuenta Cuenta) {

	index := B.buscarllave(cuenta)

	if index < B.numero && B.llaves[index].DPI == cuenta.DPI {
		if B.hoja {
			B.eliminarDeHoja(index)
		} else {
			B.eliminarDeNodo(index)
		}
	} else {
		if B.hoja {
			fmt.Println("La cuenta no existe")
			return
		}

		bandera := index == B.numero

		if B.hijos[index].numero < B.grado {
			B.llenar(index)
		}

		if bandera && index > B.numero {

			B.hijos[index - 1].eliminar(cuenta)
		} else {

			B.hijos[index].eliminar(cuenta)
		}
	}
}

// eliminarDeNodo elimina una llave de un nodo con cierto indice
func (B *BNodo) eliminarDeNodo(index int) {
	llave := B.llaves[index]

	if B.hijos[index].numero >= B.grado {
		pred := B.predecesor(index)
		B.llaves[index] = pred
		B.hijos[index].eliminar(*pred)

	} else if B.hijos[index + 1].numero >= B.grado {
		succ := B.sucesor(index)
		B.llaves[index] = succ
		B.hijos[index + 1].eliminar(*succ)

	} else {
		B.unir(index)
		B.hijos[index].eliminar(*llave)

	}
}

// eliminarDeHoja elimina la llave en cierto indice
func (B *BNodo) eliminarDeHoja(index int) {
	for i := index + 1; i < B.numero; i++ {
		B.llaves[i - 1] = B.llaves[i]
	}
	B.numero--
}

//unir uno dos nodos indicado por el indice
func (B *BNodo) unir( index int) {
	hijo := B.hijos[index]
	hermano := B.hijos[index + 1]

	hijo.llaves[B.grado- 1] = B.llaves[index]

	for i:= 0; i < hermano.numero; i++ {
		hijo.llaves[i + B.grado] = hermano.llaves[i]
	}

	if !hijo.hoja {
		for i:= 0; i <= hermano.numero; i++ {
			hijo.hijos[i + B.grado] = hermano.hijos[i]
		}
	}

	for i:= index + 1; i < B.numero; i++ {
		B.llaves[i - 1] = B.llaves[i]
	}

	for i:= index + 2; i <= B.numero; i++ {
		B.hijos[i - 1] = B.hijos[i]
	}
	hijo.numero += hermano.numero + 1
	B.numero--
}

//predecesor obtiene el anterior al indice indicado
func (B *BNodo) predecesor(index int) *Cuenta{

	aux := B.hijos[index]
	for !aux.hoja {

		aux = aux.hijos[aux.numero]
	}

	return aux.llaves[aux.numero- 1]
}

//sucesor obtiene la siguiente llave al indice
func (B *BNodo) sucesor(index int) *Cuenta{
	aux := B.hijos[index + 1]
	for !aux.hoja {
		aux = aux.hijos[0]
	}
	return aux.llaves[0]
}

// llenar llena un nodo a partir de un indice
func (B *BNodo) llenar(index int) {

	if index != 0 && B.hijos[index - 1].numero >= B.grado {

		B.prestarAnterior(index)
	} else if index != B.numero && B.hijos[index + 1].numero >= B.grado {

		B.prestarSiguiente(index)
	} else {

		if index != B.numero {
			B.unir(index)
		} else {
			B.unir(index - 1)
		}
	}
}

//prestarAnterior obtiene una llave de otro nodo
func (B *BNodo) prestarAnterior(index int) {
	hijo := B.hijos[index]
	hermano := B.hijos[index - 1]

	for i:= hijo.numero - 1; i >= 0; i-- {
		hijo.llaves[i + 1] = hijo.llaves[i]
	}

	if !hijo.hoja {
		for  i:= hijo.numero; i >= 0; i-- {
			hijo.hijos[i + 1] = hijo.hijos[i]
		}
	}

	hijo.llaves[0] = B.llaves[index - 1]

	if !hijo.hoja {
		hijo.hijos[0] = hermano.hijos[hermano.numero]
	}

	B.llaves[index - 1] = hermano.llaves[hermano.numero- 1]
	hijo.numero += 1
	hermano.numero -= 1
}

//prestarSiguiente obteine una llave del siguiente nodo
func (B *BNodo) prestarSiguiente(index int) {

	child := B.hijos[index]
	sibling := B.hijos[index + 1]
	child.llaves[(child.numero)] = B.llaves[index]

	if !child.hoja {
		child.hijos[(child.numero) + 1] = sibling.hijos[0]
	}

	B.llaves[index] = sibling.llaves[0]

	for  i:= 1; i < sibling.numero; i++ {
		sibling.llaves[i - 1] = sibling.llaves[i]
	}

	if !sibling.hoja {
		for i:= 1; i <= sibling.numero; i++ {
			sibling.hijos[i - 1] = sibling.hijos[i]
		}
	}

	child.numero += 1
	sibling.numero -= 1
}


// imprimirLlaves imprime las llaves de un nodo
func (B *BNodo) imprimirLlaves(){
	i := 0
	for i = 0; i < B.numero; i++ {
		if !B.hoja {

			B.hijos[i].imprimirLlaves()
		}

		fmt.Println(B.llaves[i])
	}
	if !B.hoja {

		B.hijos[i].imprimirLlaves()
	}
}
<<<<<<< HEAD:Backend/Archivos/BNodo.go
=======

func (B *BNodo) generarGraphviz() string {
	nombre := "nodo" + strconv.Itoa(B.llaves[0].DPI)
	texto :=  nombre + "[label = \""

	for i, llave := range B.llaves {
		if llave != nil {
			if i < B.numero {

				texto += llave.toDOT()+"|"
			} else {

				texto += llave.toDOT()
			}
		}
	}
	texto+= "\"];\n"
	for _, hijo := range B.hijos {
		if hijo != nil {
			texto += hijo.generarGraphviz()
			texto += "\t" +  nombre + "-> " + "nodo" +  strconv.Itoa(hijo.llaves[0].DPI) + "\n"
		}
	}
	return texto
}
>>>>>>> arreglo:Backend/BNodo.go

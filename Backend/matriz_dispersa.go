package main

import (
	"fmt"
	"strconv"
)

type nodoFila struct{
	dia int `json:"dia"`
	derecha *nodoFila `json:"derecha"`
	izquierda *nodoFila `json:"izquierda"`
	abajo *nodoContenido `json:"abajo"`

}

type nodoColumna struct {
	categoria string `json:"categoria"`
	arriba *nodoColumna `json:"arriba"`
	abajo *nodoColumna `json:"abajo"`
	derecha *nodoContenido `json:"derecha"`
}

type nodoContenido struct {
	contenido *ColaPedidos `json:"contenido"`
	categoria string `json:"categoria"`
	dia int `json:"dia"`
	arriba *nodoContenido `json:"arriba"`
	abajo *nodoContenido `json:"abajo"`
	izquierda *nodoContenido `json:"izquierda"`
	derecha *nodoContenido `json:"derecha"`
}

type matrizDispersa struct {
	mes string `json:"mes"`
	primeroFila *nodoFila `json:"primero_fila"`
	primeroColumna *nodoColumna `json:"primero_columna"`
}

func newMatrizDispersa(mes string) *matrizDispersa{
	return &matrizDispersa{mes, nil, nil }
}

func (matriz *matrizDispersa) agregarFila(dia int){
	temp := matriz.primeroFila

	if temp == nil { // Si el  primero de la fila es nulo agrega un nuevo nodo
		matriz.primeroFila = &nodoFila{dia, nil, nil, nil}
	} else {
		for temp.derecha != nil && temp.dia < dia {
			temp = temp.derecha
		}
		nuevo := &nodoFila{dia, nil, temp , nil}
		temp.derecha = nuevo

	}
}

func (matriz *matrizDispersa) agregarColumna(categoria string){
	temp := matriz.primeroColumna

	if temp == nil { // Si el  primero de la fila es nulo agrega un nuevo nodo
		matriz.primeroColumna = &nodoColumna{categoria, nil, nil, nil}
	} else {
		for temp.abajo != nil {
			temp = temp.abajo
		}
		nuevo := &nodoColumna{categoria, temp, nil , nil}
		temp.abajo = nuevo

	}
}

func (matriz *matrizDispersa) agregarContenido(categoria string, dia int, contenido *ColaPedidos){
	filaTemp := matriz.primeroFila
	columnaTemp := matriz.primeroColumna
	// Si no hay nada en la matriz
	if filaTemp.abajo == nil && columnaTemp.derecha == nil {
		nuevo := &nodoContenido{contenido, categoria, dia, nil, nil, nil, nil}
		matriz.primeroFila.abajo = nuevo
		matriz.primeroColumna.derecha = nuevo
		return
	}

	for filaTemp.derecha != nil && filaTemp.dia != dia {
		filaTemp = filaTemp.derecha
	}
	fmt.Print(filaTemp.dia)

	for columnaTemp.abajo != nil && columnaTemp.categoria != categoria {
		columnaTemp = columnaTemp.abajo
	}
	fmt.Println(columnaTemp.categoria)

	abajoTemp := filaTemp.abajo

	if abajoTemp != nil  {
		for abajoTemp.abajo != nil && abajoTemp.categoria != columnaTemp.categoria {
			abajoTemp = abajoTemp.abajo
		}
	}
	fmt.Print(abajoTemp)

	derTemp := columnaTemp.derecha

	if derTemp != nil {
		for derTemp.derecha != nil && derTemp.dia != filaTemp.dia {

			derTemp = derTemp.derecha
		}
	}
	fmt.Println(derTemp)
	if abajoTemp != nil && derTemp != nil { // Si ya hay nodos en la fila y columna
		nuevo := &nodoContenido{contenido, categoria, dia, abajoTemp, abajoTemp.abajo, derTemp, derTemp.derecha}
		abajoTemp.abajo = nuevo
		derTemp.derecha = nuevo
	} else if abajoTemp != nil && derTemp == nil {// Si hay nodos en la columa, pero no en la fila
		nuevo := &nodoContenido{contenido, categoria, dia, abajoTemp, abajoTemp.abajo, nil, nil}
		abajoTemp.abajo = nuevo
		columnaTemp.derecha = nuevo
	} else if abajoTemp == nil && derTemp != nil { // Si hay nodos en la fila, pero no en la columna
		nuevo := &nodoContenido{contenido, categoria, dia, nil, nil, derTemp, derTemp.derecha}
		filaTemp.abajo = nuevo
		derTemp.derecha = nuevo

	} else { // Si no hay nodos en la fila, ni en la columna
		nuevo := &nodoContenido{contenido, categoria, dia, nil, nil, nil, nil}
		filaTemp.abajo = nuevo
		columnaTemp.derecha = nuevo
	}
}

func (matriz *matrizDispersa) ordenarMatriz () {
	columnaTemp := matriz.primeroColumna

	for columnaTemp.abajo != nil  {

	}
}

func (matriz matrizDispersa) imprimirMatriz(){
	tempFila := matriz.primeroFila
	fmt.Print(" ")
	for tempFila.derecha != nil {
		fmt.Print(strconv.Itoa(tempFila.dia) + " ")
		tempFila = tempFila.derecha
	}
	fmt.Println()
	tempColumna := matriz.primeroColumna
	for tempColumna != nil {
		tempDer := tempColumna.derecha
		fmt.Print(tempColumna.categoria + " ")
		for tempDer != nil {

			fmt.Print(tempDer.contenido)
			fmt.Print(" ")
			tempDer = tempDer.derecha
		}
		fmt.Print("\n")
		tempColumna = tempColumna.abajo
	}
}

/*func main() {
	matriz := newMatrizDispersa("Enero")
	for i:=1; i< 6; i++ {
		matriz.agregarFila(i)
	}
	matriz.agregarColumna("A")
	matriz.agregarColumna("B")
	matriz.agregarColumna("C")


	temp1  := &ColaPedidos{"NA1"}
	temp2  := &ColaPedidos{"NB2"}
	temp3  := &ColaPedidos{"NC3"}
	temp4  := &ColaPedidos{"NA4"}
	temp5  := &ColaPedidos{"NB1"}
	temp6  := &ColaPedidos{"NC2"}
	temp7  := &ColaPedidos{"NA3"}
	temp8  := &ColaPedidos{"NB4"}
	temp9  := &ColaPedidos{"NC1"}
	temp10 := &ColaPedidos{"NA2"}
	temp11 := &ColaPedidos{"NB3"}
	temp12 := &ColaPedidos{"NC4"}



	matriz.agregarContenido("A", 1, temp1)
	matriz.agregarContenido("B", 2, temp2)
	matriz.agregarContenido("C", 3, temp3)
	matriz.agregarContenido("A", 4, temp4)
	matriz.agregarContenido("B", 1, temp5)
	matriz.agregarContenido("C", 2, temp6)
	matriz.agregarContenido("A", 3, temp7)
	matriz.agregarContenido("B", 4, temp8)
	matriz.agregarContenido("C", 1, temp9)
	matriz.agregarContenido("A", 2, temp10)
	matriz.agregarContenido("B", 3, temp11)
	matriz.agregarContenido("C", 4, temp12)

	matriz.imprimirMatriz()


} */

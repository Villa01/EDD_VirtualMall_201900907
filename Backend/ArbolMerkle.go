package main

import (
	_ "crypto/sha256"
	"fmt"
	"math"
	"strconv"
)


type nodoMerkle struct {
	indice   int
	altura   int
	izq, der *nodoMerkle
	hashIzq, hashDer, hash string
	info string
}

var keysha = "j1F7L_EFXcWFYlmrQvHSIMcaWxqxLjYRE04u-mJL7jw="

func newNodoMerkle(indice int) *nodoMerkle {

	return &nodoMerkle{indice, 0, nil, nil, "", "", "", ""}
}

func (n nodoMerkle) graficar(i int) string {
	var texto string

	var informacion string
	informacion += n.hash + "\n"
	if n.izq == nil && n.der == nil{

			informacion += n.info + "\\n"
	} else {
		informacion += n.hashDer
		informacion += n.hashIzq
	}

	texto+= "nodo" + strconv.Itoa(i) + "[label=\""+informacion+"\"];\n"

	return  texto
}

type arbolMerkle struct {
	raiz *nodoMerkle
	transaccion []string
}




func newArbolMerkle(transacciones []string) *arbolMerkle{
	nuevo := &arbolMerkle{nil, nil}
	for i := 0; i < calcularlongitudparahash(transacciones); i++ {
		//print(i)
		nuevo.Insertar(i)
	}

	insertarDatosUltimoNivel(&transacciones, nuevo.raiz)
	return nuevo
}

func max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func alturaMerkle(temp *nodoMerkle) int {
	if temp != nil {
		return temp.altura
	}
	return -1
}

func rotacionIzquierdaMerkle(temp **nodoMerkle) {
	aux := (*temp).izq
	(*temp).izq = aux.der
	aux.der = *temp
	(*temp).altura = max(alturaMerkle((*temp).der), alturaMerkle((*temp).izq)) + 1
	aux.altura = max(alturaMerkle((*temp).izq), (*temp).altura) + 1
	*temp = aux
}

func rotacionDerechaMerkle(temp **nodoMerkle) {
	aux := (*temp).der
	(*temp).der = aux.izq
	aux.izq = *temp
	(*temp).altura = max(alturaMerkle((*temp).der), alturaMerkle((*temp).izq)) + 1
	aux.altura = max(alturaMerkle((*temp).der), (*temp).altura) + 1
	*temp = aux
}

func rotacionDobleIzquierdaMerkle(temp **nodoMerkle) {
	rotacionDerechaMerkle(&(*temp).izq)
	rotacionIzquierdaMerkle(temp)
}

func rotacionDobleDerechaMerkle(temp **nodoMerkle) {
	rotacionIzquierdaMerkle(&(*temp).der)
	rotacionDerechaMerkle(temp)
}

func insertMerkle(indice int, root **nodoMerkle) {
	if (*root) == nil {
		*root = newNodoMerkle(indice )
		return
	}
	if indice < (*root).indice {
		insertMerkle(indice, &(*root).izq)
		if (alturaMerkle((*root).izq) - alturaMerkle((*root).der)) == -2 {
			if indice < (*root).izq.indice{
				rotacionIzquierdaMerkle(root)
			}else{
				rotacionDobleIzquierdaMerkle(root)
			}
		}
	}else if indice > (*root).indice{
		insertMerkle(indice, &(*root).der)
		if (alturaMerkle((*root).der) - alturaMerkle((*root).izq)) == 2{
			if indice > (*root).der.indice {
				rotacionDerechaMerkle(root)
			}else{
				rotacionDobleDerechaMerkle(root)
			}
		}
	}else{
	}

	(*root).altura = max(alturaMerkle((*root).izq), alturaMerkle((*root).der))+1
}

func (arbolMerkle *arbolMerkle) Insertar(indice int) {
	insertMerkle(indice, &arbolMerkle.raiz)
}

func (avl *arbolMerkle) Print(){
	inOrdenMerkle(avl.raiz)
}

func inOrdenMerkle(temp *nodoMerkle)  {
	if temp != nil {
		inOrdenMerkle(temp.izq)
		fmt.Println("Nombre: ", temp)
		inOrdenMerkle(temp.der)
	}
}

func (arbolMerkle *arbolMerkle) BuscarNodo (codigo int) * nodoMerkle{
	return buscarNodoMerkle(arbolMerkle.raiz, codigo)
}

func buscarNodoMerkle(temp *nodoMerkle, codigo int ) *nodoMerkle{
	if temp != nil {
		if temp.indice == codigo {
			return temp
		}
		izq := buscarNodoMerkle(temp.izq, codigo)
		der := buscarNodoMerkle(temp.der, codigo)

		if izq != nil {
			return izq
		} else if der != nil {
			return der
		}
	}
	return nil
}

func (a *arbolMerkle) Graficar() string {

	texto := "digraph g{\n    \n\tnode [style=\"filled\" shape=\"rectangle\" fillcolor=\"#ff00005f\"];\n"
	texto += a.graficar(a.raiz)

	texto+= "\n}"

	return  texto
}

func (a arbolMerkle) graficar(actual *nodoMerkle) string {
	var texto string

	if actual != nil{

		texto += actual.graficar(actual.indice)
		texto += a.graficar(actual.izq)
		texto += a.graficar(actual.der)

		if actual.izq != nil{
			texto+= "nodo" + strconv.Itoa(actual.indice) + " -> nodo" + strconv.Itoa(actual.izq.indice) + ";\n"
		}
		if actual.der != nil{
			texto+= "nodo" + strconv.Itoa(actual.indice) + " -> nodo" + strconv.Itoa(actual.der.indice) + ";\n"
		}
	}

	return  texto
}


func insertarDatosUltimoNivel(transacciones *[]string, temp *nodoMerkle){
	if temp != nil{
		insertarDatosUltimoNivel(transacciones,temp.izq)
		insertarDatosUltimoNivel(transacciones,temp.der)

		if temp.izq == nil && temp.der == nil{
			if len(*transacciones) != 0 {
				temp.info = (*transacciones)[0]
				temp.hash = encriptarConLlave((*transacciones)[0], keysha)
				*transacciones = append((*transacciones)[:0], (*transacciones)[1:]...) //cola

			}else{
				temp.info = ""
				temp.hash = ""
			}
		}else{
			temp.hashIzq = temp.izq.hash
			temp.hashDer = temp.der.hash
			temp.hash = encriptarConLlave(temp.hashIzq+temp.hashDer, keysha)
		}
	}

}


func calcularlongitudparahash(transacciones []string) int{
	cantidadniveles := math.Trunc(math.Log2(float64(len(transacciones))))
	devolver := 0
	for i := 0; i <= int(cantidadniveles)+1; i++ {
		p := float64(i)
		devolver += int(math.Pow(2,p))
	}
	return devolver
}


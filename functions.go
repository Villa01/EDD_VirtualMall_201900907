package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// -------------------------------Funciones del servidor ----------------------

var array []*DoublyLinkedList

// Info almacena todos los datos del json leido
var Info Information

// Index es una funcion de prueba
func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "El servidor esta funcionando")

	lista := NewDoublyLinkedList()
	store := Store{
		Name:        "Tienda de prueba",
		Description: "Descripcion de prueba",
		Contact:     "Contacto de prueba",
		Rating:      5,
	}
	store2 := Store{
		Name:        "Tienda de prueba2",
		Description: "Descripcion de prueba2",
		Contact:     "Contacto de prueba2",
		Rating:      5,
	}
	nodo := NewNode(store)
	nodo2 := NewNode(store2)
	lista.Append(nodo)
	lista.Append(nodo2)
	text, _ := lista.ToString()
	fmt.Print(text)
	num := GetAsciiValue("prueba")
	fmt.Println(num)
}

// loadStore obtiene los datos de tiendas
func loadStore(w http.ResponseWriter, req *http.Request) {

	_ = json.NewDecoder(req.Body).Decode(&Info)
	json.NewEncoder(w).Encode("Recibido")

	var matrix Matrix
	fmt.Println("$$$ LLenando la matriz...")
	matrix.fillMatrix(Info)
	fmt.Println("$$$ Matriz completamente llena")
	matrix.printMatrix()
	fmt.Println("$$$ Linealizando la matriz...")
	array = matrix.rowMajor()
	fmt.Println("$$$ Matriz completamente linealizada")
	printArray(array)
}

// getArreglo genera un reporte del vector de listas linealizado
func getArreglo(w http.ResponseWriter, req *http.Request) {

	dotArrayRoute := "reporte.dot"

	if len(array) <= 0 {
		return
	}
	text := "digraph reporte {\n"
	for i := 0; i < len(array); i++ {

		if array[i].head != nil {
			fmt.Println(array[i].head.data.Name)
			text += array[i].GetGraphviz()
		} else {
			text += "\tnode [ shape= rect label=\"Null\"] v" + fmt.Sprint(i) + ";\n"
		}
	}

	text += "\n}"
	fmt.Println(text)

	file, err := os.Create(dotArrayRoute)

	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(text)

}

// searchSpecificStore busca una tienda con los parametros que especifica el archivo json
func searchSpecificStore(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "La busqueda especifica funciona")
}

// searchByPosition busca en los arreglos las tiendas en cierta posicion
func searchByPosition(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "La busqueda por posicion funciona")
}

// deleteRegistry elimina una de las tiendas con la informacion del archivo json
func deleteRegistry(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "La eliminacion funciona")
}

// saveData guarda los datos de la matriz en un archivo json
func saveData(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "El guardado de datos funciona")
}

// --------------------- Utilidades --------------------------------------

// fillMatrix recibe la informacion y llena una matriz 3 x 3 con una lista doblemente enlazada simulando una 4ta dimension
func (matrix *Matrix) fillMatrix(info Information) *Matrix {
	var newIndexes []*IndexLetter

	data := info.Data

	// Llenar la primera dimension

	for i, dat := range data { // Recorremos cada dato obtenido del json
		newIndex := IndexLetter{
			Index: dat.Index,
		}
		newIndexes = append(newIndexes, &newIndex)
		//newIndexes[i].Index = dat.Index // Agrega la letra del indice
		var newDepartments []DepartmentMatrix

		for j, dep := range dat.Departments { // Recorremos cada departamento de cada indice
			newDepartment := DepartmentMatrix{
				name: dep.Name,
			}
			newDepartments = append(newDepartments, newDepartment)
			var newRatings [5]Rating

			for l := 1; l <= 5; l++ { // Se le crea una lista a cada posicion
				newRatings[l-1] = Rating{
					number: l,
					lista:  NewDoublyLinkedList(),
				}
			}

			for _, sto := range dep.Stores { // Recorremos cada tienda de cada departamento

				rate := int(sto.Rating) - 1
				node := NewNode(sto)
				newRatings[rate].lista.Append(node) // Se agrega la nueva tienda a la posicion del arreglo correspondiente a su calificacion
				fmt.Println("Agregué la tienda ", newRatings[rate].lista.head)
				text, _ := newRatings[rate].lista.ToString()
				fmt.Println(" a la lista ", text)
			}
			newDepartments[j].ratings = newRatings
		}
		newIndexes[i].Departments = newDepartments

	}

	matrix.indexes = newIndexes

	return matrix
}

// printMatrix imprime la matriz en un formato legible
func (matrix *Matrix) printMatrix() {

	for i := 0; i < len(matrix.indexes); i++ {
		fmt.Print(matrix.indexes[i].Index, "[ ")
		for j := 0; j < len(matrix.indexes[i].Departments); j++ {
			fmt.Print("[ ")
			for k := 0; k < len(matrix.indexes[i].Departments[j].ratings); k++ {
				fmt.Print("[ ", matrix.indexes[i].Departments[j].ratings[k].number, " ]")
			}
			fmt.Print("] ")
		}
		fmt.Println("]")
	}
}

// rowMajor linealiza la matriz a un arreglo
func (matrix *Matrix) rowMajor() []*DoublyLinkedList {
	rowSize := len(matrix.indexes)
	colSize := len(matrix.indexes[0].Departments)
	sliSize := len(matrix.indexes[0].Departments[0].ratings)
	var arrSize int = rowSize * colSize * sliSize
	var array = make([]*DoublyLinkedList, arrSize)

	for i := 0; i < rowSize; i++ {

		for j := 0; j < colSize; j++ {
			for k := 0; k < sliSize; k++ {

				array[k+sliSize*(j+colSize*i)] = matrix.indexes[i].Departments[j].ratings[k].lista
				//texto, _ := matrix.indexes[i].Departments[j].ratings[k].lista.ToString()
			}
		}
	}
	return array
}

func printArray(array []*DoublyLinkedList) {
	fmt.Print("[ ")
	for i := 0; i < len(array); i++ {
		text, _ := array[i].ToString()
		fmt.Print(" ", text, " ,")
	}
	fmt.Println("]")
}

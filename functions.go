package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// -------------------------------Funciones del servidor ----------------------

// Info almacena todos los datos del json leido
var Info Information

// Index es una funcion de prueba
func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "El servidor esta funcionando")
}

// loadStore obtiene los datos de tiendas
func loadStore(w http.ResponseWriter, req *http.Request) {

	_ = json.NewDecoder(req.Body).Decode(&Info)
	json.NewEncoder(w).Encode("Recibido")

	var matrix Matrix
	matrix.fillMatrix(Info)
	matrix.printMatrix()
	//array := matrix.rowMajor()
	//fmt.Println(len(array))
}

// getArreglo busca una tienda con los parametros que especifica el archivo json
func getArreglo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "La graficacion del arreglo funciona")
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

			for _, sto := range dep.Stores { // Recorremos cada tienda de cada departamento
				for l := 1; l <= 5; l++ { // Se le crea una lista a cada posicion
					newRatings[l-1] = Rating{
						number: l,
						lista:  NewDoublyLinkedList(),
					}
				}
				rate := int(sto.Rating) - 1
				node := NewNode(sto)
				list := *newRatings[rate].lista
				list.add(node) // Se agrega la nueva tienda a la posicion del arreglo correspondiente a su calificacion

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

	for _, index := range matrix.indexes {
		fmt.Print(index.Index, "[ ")
		for _, dep := range index.Departments {
			fmt.Print("[ ")
			for _, rat := range dep.ratings {
				fmt.Print("[ ", rat.number, " ]")
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
		fmt.Println(matrix.indexes[i].Index)

		for j := 0; j < colSize; j++ {
			fmt.Println("\t", matrix.indexes[i].Departments[j].name)
			for k := 0; k < sliSize; k++ {
				fmt.Println("\t\t", matrix.indexes[i].Departments[j].ratings[k].number)
				array[k+sliSize*(j+colSize*i)] = matrix.indexes[i].Departments[j].ratings[k].lista
				for l := 0; l < matrix.indexes[i].Departments[j].ratings[k].lista.lenght; l++ {
					node, _ := matrix.indexes[i].Departments[j].ratings[k].lista.getNodeAt(l)
					fmt.Println("\t\t", node.data.Name)
				}
			}
		}
	}

	return array
}

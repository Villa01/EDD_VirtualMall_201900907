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
func (newMatrix *Matrix) fillMatrix(info Information) *Matrix {
	fmt.Println(info.Data)
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
				fmt.Println("La calificacion de la tienda es: %d", rate)
				fmt.Println("El nodo contiene: " + node.data.Name)
				list := *newRatings[rate].lista
				texto, _ := list.toString()
				list.add(node) // Se agrega la nueva tienda a la posicion del arreglo correspondiente a su calificacion
				fmt.Println("La lista actualmente es " + texto)
			}
			newDepartments[j].ratings = newRatings
		}
		newIndexes[i].Departments = newDepartments

	}

	newMatrix.indexes = newIndexes

	return newMatrix
}

// printMatrix imprime la matriz en un formato legible
func (matrix *Matrix) printMatrix() {

	for _, index := range matrix.indexes {
		fmt.Println("Leyendo la fila ", index.Index)
		for _, dep := range index.Departments {
			fmt.Println("\tLeyendo la columna ", dep.name)
			for _, rat := range dep.ratings {
				fmt.Println("\t\tLeyendo la calificacion ", rat.number)
				texto, _ := rat.lista.toString()
				fmt.Println("\t\t\t ", texto, " y su tamaño es ", rat.lista.lenght)
			}
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// -------------------------------Funciones del servidor ----------------------

var array []VectorItem
var matrix Matrix

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

	fmt.Println("$$$ LLenando la matriz...")
	matrix.fillMatrix(Info)
	fmt.Println("$$$ Matriz completamente llena")
	//matrix.printMatrix()
	fmt.Println("$$$ Linealizando la matriz...")
	array = matrix.rowMajor()
	fmt.Println("$$$ Matriz completamente linealizada")
	//printArray(array)
}

// getArreglo genera un reporte del vector de listas linealizado
func getArreglo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("\n Generando un reporte del vector linealizado...")
	dotArrayRoute := "reporte.dot"

	if len(array) <= 0 {
		return
	}
	text := "digraph reporte {\n"
	for i := 0; i < len(array); i++ {

		if array[i].List.head != nil {
			fmt.Println(array[i].List.head.data.Name)
			text += array[i].List.GetGraphviz()
		} else {
			text += "\tnode [ shape= rect label=\"Null\"] v" + fmt.Sprint(i) + ";\n"
		}
	}

	text += "\n}"
	//fmt.Println(text)

	file, err := os.Create(dotArrayRoute)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(text)

	fmt.Println("\n EL archivo reporte.svg se encuentra en la carpeta del proyecto.")

}

// searchSpecificStore busca una tienda con los parametros que especifica el archivo json
func searchSpecificStore(w http.ResponseWriter, req *http.Request) {
	var sStore SpecificStore
	_ = json.NewDecoder(req.Body).Decode(&sStore)
	fmt.Println("$$$Buscando tienda con los parametros especificados")
	var store Store
	for i := 0; i < len(array); i++ {
		if array[i].Department == sStore.Departament && array[i].Rating == sStore.Rating {
			for j := 0; j < array[i].List.lenght; j++ {
				tempNode, _ := array[i].List.GetNodeAt(j)
				tempName := tempNode.data.Name
				//fmt.Println(tempName)
				if tempName == sStore.Name {
					store = tempNode.data
				}
			}
		}
	}

	fmt.Println("$$$ Retornando datos obtenidos")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(store)
}

// searchByPosition busca en los arreglos las tiendas en cierta posicion
func searchByPosition(w http.ResponseWriter, req *http.Request) {

	if len(array) == 0 {
		fmt.Println("$$$Primero debe llenar el arreglo con informacion")
		json.NewEncoder(w).Encode("Primero debe llenar el arreglo con informacion")
		return
	}

	parameters := mux.Vars(req)
	index, err := strconv.Atoi(parameters["numero"])

	if err != nil {
		panic(err)
	}

	fmt.Println("$$$Iniciando la busqueda del elemento en el arreglo linealizado")

	item := array[index]
	nodes := item.List.GetJSONNodes()
	if len(nodes.Nodes) > 0 {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(nodes)
	} else {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode("La posicion está vacía.")
	}

}

// deleteRegistry elimina una de las tiendas con la informacion del archivo json
func deleteRegistry(w http.ResponseWriter, req *http.Request) {
	var sStore SpecificStore
	_ = json.NewDecoder(req.Body).Decode(&sStore)

	if len(array) == 0 {
		fmt.Println("$$$Primero debe llenar el arreglo con informacion")
		json.NewEncoder(w).Encode("Primero debe llenar el arreglo con informacion")
		return
	}

	fmt.Println("$$$Buscando tienda con los parametros especificados")
	for i := 0; i < len(array); i++ {
		if array[i].Department == sStore.Departament && array[i].Rating == sStore.Rating {
			for j := 0; j < array[i].List.lenght; j++ {
				tempNode, _ := array[i].List.GetNodeAt(j)
				tempName := tempNode.data.Name
				if tempName == sStore.Name {
					array[i].List.DeleteNode(j)
				}
			}
		}
	}
	printArray(array)

}

// saveData guarda los datos de la matriz en un archivo json
func saveData(w http.ResponseWriter, req *http.Request) {
	m := inverseRowMajor(array)
	m.printMatrix()
	enableCors(&w)
	json.NewEncoder(w).Encode(m)
}

func getTiendas(w http.ResponseWriter, req *http.Request){
	fmt.Println("$$$Devolviendo las tiendas")
	stores := fillStores(array)
	//fmt.Print(stores)
	enableCors(&w)
	json.NewEncoder(w).Encode(stores)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
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
				Name: dep.Name,
			}
			newDepartments = append(newDepartments, newDepartment)
			var newRatings [5]Rating

			for l := 1; l <= 5; l++ { // Se le crea una lista a cada posicion
				newRatings[l-1] = Rating{
					Number: l,
					Lista:  NewDoublyLinkedList(),
				}
			}

			for _, sto := range dep.Stores { // Recorremos cada tienda de cada departamento

				rate := int(sto.Rating) - 1
				node := NewNode(sto)
				newRatings[rate].Lista.Append(node) // Se agrega la nueva tienda a la posicion del arreglo correspondiente a su calificacion
			}
			newDepartments[j].Ratings = newRatings
		}
		newIndexes[i].Departments = newDepartments

	}

	matrix.Indexes = newIndexes

	return matrix
}


func fillStores(vector []VectorItem) []Store{
	var stores []Store
	for _, item := range vector {
		for i:= 0; i < item.List.lenght; i++{
			temp,_ := item.List.GetNodeAt(i)
			stores = append(stores, temp.data)
		}
	}

	return stores
}

// printMatrix imprime la matriz en un formato legible
func (matrix *Matrix) printMatrix() {

	for i := 0; i < len(matrix.Indexes); i++ {
		fmt.Print(matrix.Indexes[i].Index, "[ ")
		for j := 0; j < len(matrix.Indexes[i].Departments); j++ {
			fmt.Print("[ ")
			for k := 0; k < len(matrix.Indexes[i].Departments[j].Ratings); k++ {
				text, _ := matrix.Indexes[i].Departments[j].Ratings[k].Lista.ToString()
				fmt.Print("[ ", text, " ]")

			}
			fmt.Print("] ")
		}
		fmt.Println("]")
	}
}

// rowMajor linealiza la matriz a un arreglo
func (matrix *Matrix) rowMajor() []VectorItem {
	rowSize := len(matrix.Indexes)
	colSize := len(matrix.Indexes[0].Departments)
	sliSize := len(matrix.Indexes[0].Departments[0].Ratings)
	var arrSize int = rowSize * colSize * sliSize
	var array = make([]VectorItem, arrSize)

	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			for k := 0; k < sliSize; k++ {
				department := matrix.Indexes[i].Departments[j].Name
				rating := matrix.Indexes[i].Departments[j].Ratings[k].Number
				list := matrix.Indexes[i].Departments[j].Ratings[k].Lista

				temp := VectorItem{
					Department: department,
					Rating:     rating,
					List:       list,
				}
				array[k+sliSize*(j+colSize*i)] = temp
				//texto, _ := matrix.indexes[i].Departments[j].ratings[k].lista.ToString()
			}
		}
	}
	return array
}

func inverseRowMajor(array []VectorItem) Matrix {

	rowSize := len(matrix.Indexes)
	colSize := len(matrix.Indexes[0].Departments)
	sliSize := len(matrix.Indexes[0].Departments[0].Ratings)

	var indexes []*IndexLetter
	var matrix Matrix

	for i := 0; i < rowSize; i++ {
		var newIndex IndexLetter
		indexes = append(indexes, &newIndex)
		var departs []DepartmentMatrix
		for j := 0; j < colSize; j++ {
			var dep DepartmentMatrix
			departs = append(departs, dep)

			var ratings [5]Rating
			for k := 0; k < 5; k++ {
				var rat Rating
				ratings[k] = rat
			}
			departs[j].Ratings = ratings
		}
		indexes[i].Departments = departs
	}
	matrix.Indexes = indexes

	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			for k := 0; k < sliSize; k++ {
				matrix.Indexes[i].Departments[j].Ratings[k].Lista = array[k+sliSize*(j+colSize*i)].List
			}
		}
	}
	return matrix
}

func printArray(array []VectorItem) {
	fmt.Print("[ ")
	for i := 0; i < len(array); i++ {
		text, _ := array[i].List.ToString()
		fmt.Print(" ", text, " ,")
	}
	fmt.Println("]")
}

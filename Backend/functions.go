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
var carrito []Product

// Info almacena todos los datos del json leido
var Info Information
var InvResp InventoryResponse

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
	enableCors(&w)
}

func loadInventories(w http.ResponseWriter, req *http.Request) {

	err := json.NewDecoder(req.Body).Decode(&InvResp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - La información no es correcta"))
		return
	}
	enableCors(&w)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("200 - La información fue recibida"))
	json.NewEncoder(w).Encode("Recibido")

	fmt.Println("$$$ Asignando el inventario a cada tienda")
	asignInventories()
	fmt.Println("$$$ Inventarios asignados")
	//matrix.printMatrix()
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
	fmt.Println("$$$ Devolviendo las tiendas")
	stores := fillStores(array)
	//fmt.Print(stores)
	enableCors(&w)
	json.NewEncoder(w).Encode(stores)
}

func getProductos(w http.ResponseWriter, req *http.Request){
	fmt.Println("$$$ Devolviendo los productos")
	productos := obtainProducts()
	//fmt.Print(stores)
	enableCors(&w)
	json.NewEncoder(w).Encode(productos)
}

func agregarAlCarrito(w http.ResponseWriter, req *http.Request){

	enableCors(&w)
	var producto Product
	_ = json.NewDecoder(req.Body).Decode(&producto)
	fmt.Println("$$$ Producto " + producto.Nombre + " agregado al carrito de compras")
	invProducto := searchProduct(producto)
	if invProducto == nil {
		fmt.Println("$$$ 409 - no se encontró el artículo en el inventario")
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("No se encontró el artículo en el inventario"))
		return
	}

	if invProducto.Cantidad > 0 && invProducto.Cantidad >= producto.Cantidad{
		//zinvProducto.Cantidad -= producto.Cantidad
		carrito = append(carrito, producto)
	} else {
		fmt.Println("$$$ 409 - El artículo no cuenta con inventario")
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("El artículo no cuenta con inventario"))

	}

}


func obtenerCarrito(w http.ResponseWriter, req *http.Request){
	fmt.Println("$$$ Devolviendo el carrito")
	enableCors(&w)
	json.NewEncoder(w).Encode(carrito)
}

func eliminarDelCarrito(w http.ResponseWriter, req *http.Request){
	fmt.Println("$$$ Eliminando del carrito")
	enableCors(&w)
	var eliminado Product
	json.NewDecoder(req.Body).Decode(&eliminado)

	for i, product := range carrito {
		if product.Codigo == eliminado.Codigo {
			carrito = append(carrito[:i], carrito[i+1:]...)
		}
	}
}

func hacerPedido( w http.ResponseWriter, req *http.Request){
	fmt.Println("%%% Haciendo nuevo pedido")
	enableCors(&w)
	var pedido []Product
	json.NewDecoder(req.Body).Decode(&pedido)

	for _, product := range pedido {
		if !verificarExistencias(product){
			fmt.Println("$$$ 409 - El artículo "+product.Nombre+" no cuenta con inventario")
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("El artículo no cuenta con inventario"))

			return
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// --------------------- Utilidades --------------------------------------

func verificarExistencias(producto Product) bool{
	temp := searchProduct(producto)
	if producto.Cantidad > temp.Cantidad {
		return false
	}
	return true
}

func searchProduct(producto Product) *Product{
	var retorno *Product

	for _, item := range array {
		for i := 0; i< item.List.lenght; i++ {
			node,_ := item.List.GetNodeAt(i)
			if node.data.Inventory != nil{
				temp := node.data.Inventory.BuscarNodo(producto.Codigo)

				if temp != nil{
					retorno = &(temp.producto)
				}
			}
		}
	}
	return retorno
}

func asignInventories(){

	for _, inventario := range InvResp.Invetarios {
		name := inventario.Tienda
		depart := inventario.Departamento
		rat := inventario.Calificacion
		// Se recorren los items del arreglo linealizado para encontrar la tienda
		for _, item := range array {
			// Cuando se encuentra la tienda

			if item.Department == depart && item.Rating == rat {
				temp := item.List.searchByContent(name)
				if temp == nil {
					break
				}
				if temp.data.Inventory == nil {
					temp.data.Inventory = NewAVL()
				}
				// Se le asignan todos los productos a su inventario
				for _, producto := range inventario.Productos {
					// Se verifica si el indice ya existe
					existente := temp.data.Inventory.BuscarNodo(producto.Codigo)
					if existente != nil {

						existente.producto.Cantidad += producto.Cantidad
					} else {
						// Si no existe
						temp.data.Inventory.Insertar(producto, producto.Codigo)
					}

				}
			}
		}

	}
}


func obtainProducts() []Product{
	var products []Product
	for _, item := range array {
		for i := 0; i< item.List.lenght; i++ {
			node,_ := item.List.GetNodeAt(i)
			if node.data.Inventory != nil {

				tempProd := node.data.Inventory.ObtenerProductos()
				if tempProd != nil {

					products = append(products,tempProd...)
				}
			}
		}
	}
	return products
}



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

package main

import (
	"fmt"
	"net/http"
)

// Index es una funcion de prueba
func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "El servidor esta funcionando")
}

// loadStore obtiene los datos de tiendas
func loadStore(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "La carga de tiendas esta conectada")
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

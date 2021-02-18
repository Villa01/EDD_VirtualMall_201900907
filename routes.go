package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route permite crear rutas con todas sus propiedades
type Route struct {
	Name     string
	Method   string
	Path     string
	Function http.HandlerFunc
}

// Routes es un arrego de Route
type Routes []Route

// Aqui se agregan las nuevas rutas
var routes = Routes{
	Route{Name: "Index", Method: "GET", Path: "/", Function: Index},
	Route{Name: "CargaTiendas", Method: "POST", Path: "/guardar", Function: loadStore},
	Route{Name: "GetArreglo", Method: "GET", Path: "/getArreglo", Function: getArreglo},
	Route{Name: "TiendaEspecifica", Method: "POST", Path: "/TiendaEspecifica", Function: searchSpecificStore},
	Route{Name: "BusquedaPorPosicion", Method: "GET", Path: "/id/numero,", Function: searchByPosition},
	Route{Name: "EliminarRegistro", Method: "DELETE", Path: "/Eliminar", Function: deleteRegistry},
	Route{Name: "GuardarDatos", Method: "GET", Path: "/Guardar", Function: saveData}}

// CreateRouter construye un router con todas las rutas en routes
func CreateRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		// Agrega todos los atributos al router
		router.Name(route.Name).Methods(route.Method).Path(route.Path).Handler(route.Function)

	}
	return router
}

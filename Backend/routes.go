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
	Route{Name: "CargaTiendas", Method: "POST", Path: "/cargarTienda", Function: loadStore},
	Route{Name: "CargarInventarios", Method: "POST", Path: "/cargarInventarios", Function: loadInventories},
	Route{Name: "GetProductos", Method: "GET", Path: "/getProductos", Function: getProductos},
	Route{Name: "GetArreglo", Method: "GET", Path: "/getArreglo", Function: getArreglo},
	Route{Name: "TiendaEspecifica", Method: "POST", Path: "/TiendaEspecifica", Function: searchSpecificStore},
	Route{Name: "BusquedaPorPosicion", Method: "GET", Path: "/id/{numero}", Function: searchByPosition},
	Route{Name: "EliminarRegistro", Method: "DELETE", Path: "/Eliminar", Function: deleteRegistry},
	Route{Name: "GuardarDatos", Method: "GET", Path: "/Guardar", Function: saveData},
	Route{Name: "getTiendas", Method: "GET", Path: "/getTiendas", Function: getTiendas},
	Route{Name: "agregarAlCarrito", Method: "POST", Path: "/agregarAlCarrito", Function: agregarAlCarrito},
	Route{Name: "getCarrito", Method: "GET", Path: "/getCarrito", Function: obtenerCarrito},
	Route{Name: "eliminarDelCarrito", Method: "POST", Path: "/eliminarDelCarrito", Function: eliminarDelCarrito},
	Route{Name: "hacerPedido", Method: "POST", Path: "/hacerPedido", Function: hacerPedido},
	Route{Name: "cargarPedidos", Method: "POST", Path: "/cargarPedidos", Function: cargarPedidos},
	Route{Name: "verificacionLogIn", Method: "POST", Path: "/verificacionLogIn", Function: verificarPassword},
	Route{Name: "obtenerCuentaActual", Method: "GET", Path: "/cuentaActual", Function: obtenerCuenta},
}
// CreateRouter construye un router con todas las rutas en routes
func CreateRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		// Agrega todos los atributos al router
		router.Name(route.Name).Methods(route.Method).Path(route.Path).Handler(route.Function)

	}
	return router
}

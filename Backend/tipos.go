package main

import "strconv"

// DoublyLinkedList es una lista doble enlazada
type DoublyLinkedList struct {
	head   *Node
	lenght int
}

// NewDoublyLinkedList crea una DoublyLinkedList vacia
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, 0}
}

// Node es un nodo con una tienda dentro
type Node struct {
	data     Store
	next     *Node
	previous *Node
}

// Store es un tipo donde se puede almacenar toda la información de una tienda
type Store struct {
	Name        string `json:"Nombre"`
	Description string `json:"Descripcion"`
	Contact     string `json:"Contacto"`
	Rating      int8   `json:"Calificacion"`
	Logo        string `json: Logo`
	Inventory   *AVL
}

// Department se refiere a los departamentos a los que pertenencen las tiendas
type Department struct {
	Name   string  `json:"Nombre"`
	Stores []Store `json:"Tiendas"`
}

// Datum es una matriz 3 x 3 que contiene indices y departamentos
type Datum struct {
	Index       string       `json:"Indice"`
	Departments []Department `json:"Departamentos"`
}

// Information es un tipo que contiene informacion de las tienas de un centro comercial
type Information struct {
	Data []Datum `json:"Datos"`
}

//IndexLetter es el tipo de dato de las filas de la matriz
type IndexLetter struct {
	Index       string
	Departments []DepartmentMatrix `json: "Departamentos`
}

//DepartmentMatrix es el tipo de dato de las columnas de la matriz
type DepartmentMatrix struct {
	Name    string    `json:"Departamento"`
	Ratings [5]Rating `json:"Calificaciones"`
}

// Rating es la tercera dimension de la matriz
type Rating struct {
	Number int `json: "Calificacion"`
	Lista  *DoublyLinkedList
}

// Matrix es una matriz llenada con informacion de un json
type Matrix struct {
	Indexes []*IndexLetter
}

// JSONNodes es una representacion de una store compatible con json
type JSONNodes struct {
	Name        string `json:"Nombre"`
	Description string `json:"Descripcion"`
	Contact     string `json:"Contacto"`
}

// SeveralJSONNodes es un conjunto de JSONNodes
type SeveralJSONNodes struct {
	Nodes []JSONNodes `json:"Tiendas"`
}

// SpecificStore es un tipo que permite buscar por medio de 3 parametros
type SpecificStore struct {
	Departament string `json:"Departamento"`
	Name        string `json:"Nombre"`
	Rating      int    `json:"Calificacion"`
	Logo        string `json: Logo`
}

// VectorItem es un item para llenar el vector linealizado
type VectorItem struct {
	Department string `json:"Departamento"`
	Rating     int    `json:"Calificacion"`
	List       *DoublyLinkedList
}

// Product son los productos del inventario
type Product struct {
	Nombre         string  `json: "Nombre"`
	Codigo         int     `json: "Codigo"`
	Descripcion    string  `json: "Descripcion"`
	Precio         float32 `json: "Precio"`
	Cantidad       int     `json: "Cantidad"`
	Imagen         string  `json:"Imagen"`
	Almacenamiento string  `json:"Almacenamiento"`
}

type Invetory struct {
	Tienda       string    `json:"Tienda"`
	Departamento string    `json:"Departamento"`
	Calificacion int       `json:"Calificacion"`
	Productos    []Product `json:"Productos"`
}

type InventoryResponse struct {
	Invetarios []Invetory `json:"Inventarios"`
}

type ColaPedidos struct {
	mes     int      `json:"mes"`
	dia     int      `json:"dia"`
	pedidos []Pedido `json:"pedidos"`
}

func nuevaCola(mes int, dia int) *ColaPedidos {
	var pedidos []Pedido
	return &ColaPedidos{mes, dia, pedidos}
}

type Pedido struct {
	Fecha        string    `json:"Fecha"`
	Tienda       string    `json:"Tienda"`
	Departamento string    `json:"Tienda"`
	Calificacion int       `json:"Calificacion"`
	Productos    []Product `json:"Productos"`
}

type PedidosResponse struct {
	Pedidos []Pedido `json:"Pedidos"`
}

type VerificacionLogInResponse struct {
	DPI      int    `json:"DPI"`
	Password string `json:"password"`
}

type RespuestaVerificacionPassword struct {
	Correcta bool   `json:"correcta"`
	Cuenta   Cuenta `json:"cuenta"`
}

type Cuenta struct {
	DPI    int    `json:"Dpi"`
	Nombre string `json:"Nombre"`
	Email  string `json:"Correo"`
	Contra string `json:"Password"`
	Cuenta string `json:"Usuario"`
}

type EliminarResponse struct {
	Password string `json:"password"`
}

type CargaUsuarios struct {
	Cuentas []Cuenta `json:"Usuarios"`
}

type RespuestaBooleana struct {
	booleano bool `json:"booleano"`
}

type RespuestaString struct{
	Texto string `json:"texto"`
}
type RespuestaReportes struct {
	Reportes []Reporte `json:"Reportes"`
}

type Reporte struct {
	Nombre string `json:"nombre"`
	Ruta   string `json:"ruta"`
}

func (c Cuenta) toDOT() string {
	return strconv.Itoa(c.DPI) + "\\n" + c.Nombre + "\\n" + c.Email + "\\n" + c.Contra + "\\n" + c.Cuenta
}

func (c Cuenta) toDotEncriptado() string {
	return encriptarConLlave(strconv.Itoa(c.DPI),llave) + "\\n" + encriptarConLlave(c.Nombre,llave) + "\\n" + encriptarConLlave(c.Email,llave) + "\\n" +encriptarConLlave(c.Contra,llave) + "\\n" +encriptarConLlave(c.Cuenta,llave)
}

func (c Cuenta) toStringEncripSensible() string {
	return encriptarConLlave(strconv.Itoa(c.DPI),llave) + "\\n" + c.Nombre + "\\n" + encriptarConLlave(c.Email,llave) + "\\n" + encriptarConLlave(c.Contra,llave) + "\\n" + c.Cuenta
}


type GrafoResponse struct {
	Nodos                []Nodo `json:"Nodos"`
	PosicionInicialRobot string `json:"PosicionInicialRobot"`
	Entrega              string `json:"Entrega"`
}

type Nodo struct {
	Nombre  string   `json:"Nombre"`
	Enlaces []Enlace `json:"Enlaces"`
}

type Enlace struct {
	Nombre    string `json:"Nombre"`
	Distancia int  `json:"Distancia"`
}


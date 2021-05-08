package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)
var rutaReportesDot string
var rutaReportesPng string
var timer int
var ArbolUsuarios arbolMerkle
var ArbolTiendas arbolMerkle
var ArbolProductos arbolMerkle
var ArbolTransacciones arbolMerkle

var usuariosMer []string
var tiendasMer []string
var productosMer []string
var transaccionesMer []string

func main() {
	obtenerUltimoHash()
	admin := &Cuenta{
		DPI:    1234567890101,
		Nombre: "EDD2021",
		Email:  " auxiliar@edd.com",
		Contra: "1234",
		Cuenta: "Admin",
	}
	encriptar("Imprime una llave")

	timer =5000
	/*
	var info []string
	info = append(info, "Info1")
	info = append(info, "Info2")
	info = append(info, "Info3")
	info = append(info, "Info4")

	a := newArbolMerkle(info)

	for i := 0; i < 15; i++ {
		a.Insertar(i)
	}
	*/
	//b1 := NewBloque(0, 0000, "05-08-2021::12:32:05", "5WePmdCjmEOI471dpwC2574HgWXk5rhTBYOsR5AQc-0=", "0000", "")
	//b1.crearHash()
	//b1.guardarBloque()

	//escribirDOT(a.Graficar(), "arbolMerkle")

	rutaReportesDot = "C:\\Users\\javil\\go\\src\\Proyecto 3\\EDD_VirtualMall_201900907\\Backend\\reportesDot"
	rutaReportesPng = "C:\\Users\\javil\\go\\src\\Proyecto 3\\EDD_VirtualMall_201900907\\Frontend\\VirtualMall\\src\\assets"

	ArbolCuentas = *NuevoBTree(5)
	ArbolCuentas.Insert(*admin)

	//port := ":3000"
	router := CreateRouter()
	//log.Fatal(http.ListenAndServe(port, router))
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))


}

func Cada5Min(){
	NewTimer(timer, func() {
		fmt.Println("Creando nuevo bloque\n")
		fecha := obtenerFecha()
		data := obtenerDatos()
		indice := obtenerIndice()
		var previous string

		if indice ==0 {
			previous = "0000"
		} else {
			previous = obtenerUltimoHash()
		}
		b1 := NewBloque(indice, 0000, fecha, data, previous, "")
		b1.guardarBloque()

		ArbolProductos = *newArbolMerkle(productosMer)
		ArbolUsuarios = *newArbolMerkle(usuariosMer)
		ArbolTiendas = *newArbolMerkle(tiendasMer)
		ArbolTransacciones = *newArbolMerkle(transaccionesMer)


		Cada5Min()
	})
}


func NewTimer(seconds int, action func()) *time.Timer {
	timer := time.NewTimer(time.Second * time.Duration(seconds))

	go func() {
		<-timer.C
		action()
	}()

	return timer
}

func leerBloques(){
	fptr := flag.String("fpath", "Bloques/0.txt", "direccion")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func obtenerFecha() string{
	ahora := time.Now()

	return ahora.Format("02-01-06::15:04:05")

}

func obtenerDatos() string {
	var datos string

	if ArbolProductos.raiz != nil{
		datos += ArbolProductos.raiz.hash + "\\n"
	}
	if ArbolUsuarios.raiz != nil{
		datos += ArbolUsuarios.raiz.hash + "\\n"
	}
	if ArbolTiendas.raiz != nil{
		datos += ArbolTiendas.raiz.hash + "\\n"
	}
	if ArbolTransacciones.raiz != nil{
		datos += ArbolTransacciones.raiz.hash + "\\n"
	}

	return datos
}

func obtenerIndice() int {
	files,err := ioutil.ReadDir("Bloques/")
	if(err != nil){
		fmt.Println("$$$ Error al tratar de obtener el indice del bloque")
	}
	fmt.Println(len(files))
	return len(files)
}

func obtenerUltimoHash() string{
	indice := obtenerIndice()
	fptr := flag.String("fpath", "Bloques/"+strconv.Itoa(indice-1)+".txt", "direccion")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	contador := 0
	for s.Scan() {
		if contador == 5{
			fmt.Println(s.Text())
			return s.Text()
		}
		contador++
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	return "Error"
}
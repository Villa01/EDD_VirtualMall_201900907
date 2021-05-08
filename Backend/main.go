package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)
var rutaReportesDot string
var rutaReportesPng string
func main() {

	admin := &Cuenta{
		DPI:    1234567890101,
		Nombre: "EDD2021",
		Email:  " auxiliar@edd.com",
		Contra: "1234",
		Cuenta: "Admin",
	}
	encriptar("Imprime una llave")

	sum := sha256.Sum256([]byte("hello world\n"))
	sum2 := sha256.Sum256([]byte("hello worldaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n"))
	fmt.Printf("%x", sum)
	fmt.Printf("%x", sum2)

	var info []string
	info = append(info, "Info1")
	info = append(info, "Info2")
	info = append(info, "Info3")
	info = append(info, "Info4")

	a := newArbolMerkle(info)

	for i := 0; i < 15; i++ {
		a.Insertar(i)
	}

	escribirDOT(a.Graficar(), "arbolMerkle")

	rutaReportesDot = "C:\\Users\\javil\\go\\src\\Proyecto 3\\EDD_VirtualMall_201900907\\Backend\\reportesDot"
	rutaReportesPng = "C:\\Users\\javil\\go\\src\\Proyecto 3\\EDD_VirtualMall_201900907\\Frontend\\VirtualMall\\src\\assets"

	ArbolCuentas = *NuevoBTree(5)
	ArbolCuentas.Insert(*admin)

	//port := ":3000"
	router := CreateRouter()
	//log.Fatal(http.ListenAndServe(port, router))
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))


}

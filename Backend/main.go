package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func main() {

	admin := &Cuenta{
		DPI:    1234567890101,
		Nombre: "EDD2021",
		Email:  " auxiliar@edd.com",
		Contra: "1234",
		Cuenta: "Admin",
	}



	ArbolCuentas = *NuevoBTree(5)
	ArbolCuentas.Insert(*admin)
	//port := ":3000"
	router := CreateRouter()
	//log.Fatal(http.ListenAndServe(port, router))
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))


}

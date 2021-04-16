package main

import (
	"fmt"
	"strconv"
)

func main() {
	json := "{\n    \"Usuarios\": [\n"
	dpi := 1234567890102
	for i := 0; i < 100; i++ {
		json += "\n\t{\n\t\t\"Dpi\": " + strconv.Itoa(dpi) + "," +
			"\n\t\t\"Nombre\": \"Nombre" + strconv.Itoa(i) + "\", " +
			"\n\t\t\"Correo\": \"Correo" + strconv.Itoa(i) + "@correo.com\"," +
			"\n\t\t\"Password\": \"hola" + strconv.Itoa(i) + "\"," +
			"\n\t\t\"Cuenta\": \"Usuario\"" +
			"\n\t},"

		dpi++
	}
	json += "\n}"
	fmt.Println(json)
}

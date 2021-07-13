package main

import (
	"fmt"

	"rsc.io/quote" // Aqui usamos un paquete que fue creado por otro programador go, La funcion de este paquete es buscar una frase inspiradora y cargarla como variable
	//(Ausumo que trae la frase en el paquete) ese paquete se instala con el comando go mod tidy
)

func main() {
	fmt.Println(quote.Go())
}

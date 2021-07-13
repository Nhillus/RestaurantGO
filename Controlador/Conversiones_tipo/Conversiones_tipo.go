package main

import (
	"fmt" // asi se puede traer varios paquetes actualizar con go mod tidy siempre que traigas un paquete
	"strconv"
)

func main() {
	edad := 29

	edad_str := strconv.Itoa(edad) // este metodo permite tranformar de string a entero Atoi y entero a string Itoa, en esta ocasion use Itoa

	fmt.Println("Mi edad es", edad)
	fmt.Println("Mi edad en string es " + edad_str)

}

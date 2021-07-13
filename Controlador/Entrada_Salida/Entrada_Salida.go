package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	edad := 29
	fmt.Printf("Mi edad es %d\n", edad) //Esta es una forma muy parecida a imprimir en c o c++ la d es base 10
	Nombre := "Thor"
	fmt.Printf("El nombre del dios del trueno es = %s\n", Nombre) // el verbo s es para un string o slice
	Flag := true
	fmt.Printf("El valor de la flag es: %t\n", Flag) // el t es para valores true o false
	precio := 20.5
	fmt.Printf("El valor es: %f\n", precio)                        //f para decimales
	fmt.Printf("El valor en valores cientificos es: %e\n", precio) //e para cientifico
	var tu_edad int
	fmt.Println("Coloca tu edad")
	fmt.Scanf("%d\n", &tu_edad)
	fmt.Printf("la edad de esta persona es: %d\n", tu_edad)

	reader := bufio.NewReader(os.Stdin) // revisar este codigo
	fmt.Println("Ingresa tu nombre:  ")
	nombre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre)
	}

}

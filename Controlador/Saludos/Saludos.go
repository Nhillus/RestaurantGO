package main

import (
	"bufio"
	"fmt"
	"os"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Cual es tu nombre?")
	nombre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(Hello(nombre)) // Sinceramente no me gusta esta forma de realizar esta impresion
	}
}

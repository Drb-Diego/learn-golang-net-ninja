package main

import "fmt"

func main() {

	// Declarando o tipo da variavel
	var name string = "Hello world!"

	// Inferindo o tupo da variavel
	var nameTwo = "Hello world!"

	// Declarando variaveis sem inicializar
	var nameThree string

	fmt.Println(name, nameTwo)

	nameThree = "Hello world!"
	fmt.Println(name, nameTwo, nameThree)

	nameFour := "something awesome"
	fmt.Println(nameFour)
}

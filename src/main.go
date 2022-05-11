package main

import "fmt"

func main() {

	helloMessage := "Hello"
	worldMessage := "World"

	// Println - Print and newline
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf
	text := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(text)

	// Type of variable
	fmt.Printf("%T\n", nombre)
	fmt.Printf("%T\n", cursos)

}

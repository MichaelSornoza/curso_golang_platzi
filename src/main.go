package main

import "fmt"

func main() {

	// Declare a constant
	fmt.Println("Constants----------------------------------------------------")

	const Pi float64 = 3.14
	const Pi2 = 3.14
	const text string = "Hello World"

	fmt.Println("Pi: ", Pi)
	fmt.Println("Pi2: ", Pi2)
	fmt.Println("text: ", text)

	// Declare a variable
	fmt.Println("Variables----------------------------------------------------")

	base := 12
	var height int = 14
	var area int = base * height

	fmt.Println("base: ", base)
	fmt.Println("height: ", height)
	fmt.Println("area: ", area)

	// Zero values
	fmt.Println("Zero values----------------------------------------------------")

	var zero int
	var zerob float64
	var zerostring string
	var zerobool bool

	fmt.Println("zero: ", zero)
	fmt.Println("zerob: ", zerob)
	fmt.Println("zerostring: ", zerostring)
	fmt.Println("zerobool: ", zerobool)

	zero = 12
	zerob = 12.34
	zerostring = "Hello World"
	zerobool = true

	fmt.Println("zero: ", zero)
	fmt.Println("zerob: ", zerob)
	fmt.Println("zerostring: ", zerostring)
	fmt.Println("zerobool: ", zerobool)

	// Excercise 1
	fmt.Println("Excercise 1----------------------------------------------------")

	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("squareArea: ", squareArea)
}

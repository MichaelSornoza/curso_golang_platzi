package main

import "fmt"

func main() {

	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("The square Area is: ", squareArea)

	x := 30
	y := 10

	result := x + y
	fmt.Println("The result is: ", result)

	result = x - y
	fmt.Println("The result of the subtraction is: ", result)

	result = x * y
	fmt.Println("The result of the multiplication is: ", result)

	result = x / y
	fmt.Println("The result of the division is: ", result)

	result = x % y
	fmt.Println("The result of the rest is: ", result)

	result = x ^ y
	fmt.Println("The result of the exponentiation is: ", result)

	result = x & y
	fmt.Println("The result of the bitwise AND is: ", result)

	result = x | y
	fmt.Println("The result of the bitwise OR is: ", result)

	result++
	fmt.Println("The result of the increment is: ", result)

	result--
	fmt.Println("The result of the decrement is: ", result)

	result += y
	fmt.Println("The result of the addition is: ", result)

	result -= y
	fmt.Println("The result of the subtraction is: ", result)

}

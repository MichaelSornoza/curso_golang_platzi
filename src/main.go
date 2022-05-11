package main

import "fmt"

func printAMessage(message string) {

	fmt.Println(message)

}

func printAMessageWithThreeArgs(message, sender, receiver string) {

	finalMessage := fmt.Sprintf("%s: -> %s \nReceived: %s", sender, message, receiver)

	fmt.Println(finalMessage)
}

func returnValue(value int) int {
	return value * 2
}

func doubleReturn(value int) (int, int) {
	return value * 2, value * 3
}

func main() {

	printAMessage("Hello World")

	printAMessageWithThreeArgs("Hello World", "John", "Mary")

	value := returnValue(2)
	fmt.Println(value)

	value1, value2 := doubleReturn(2)
	fmt.Println(value1, value2)

	value3, _ := doubleReturn(3)
	fmt.Println(value3)
}

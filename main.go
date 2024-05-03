package main

import (
	"fmt"
)

func sayHello(name string, index int) {
	for i := 0; i < index; i++ {
		fmt.Printf("Hi %s\n", name)
	}
}
func math(x int, y int) int {
	fmt.Println("first number is ", x)
	fmt.Println("second number is ", y)
	fmt.Println("now do the math with", x, "+", y)
	return x + y
}
func main() {
	sayHello("Amp", 2)
	number1 := 1
	number2 := 2

	allNumber := math(number1, number2)
	fmt.Println("Sum =", allNumber)
}

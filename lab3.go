package main

import (
	"pplabi/mathutils"
	"pplabi/stringutils"
	"fmt"
)

func main() {
	// Факториал числа
	var num int
	fmt.Print("Enter a number to calculate factorial: ")
	fmt.Scan(&num)
	fmt.Printf("Factorial of %d is %d\n", num, mathutils.Factorial(num))

	// Переворот строки
	var inputString string
	fmt.Print("Enter a string to reverse: ")
	fmt.Scan(&inputString)
	fmt.Printf("Reversed string: %s\n", stringutils.Reverse(inputString))

}

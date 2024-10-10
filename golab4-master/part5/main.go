package main

import (
	"fmt"
)

func main() {
	var n int
	var sum int

	fmt.Println("Введите количество чисел:")
	fmt.Scan(&n)

	fmt.Println("Введите числа по одному:")

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		sum += num
	}

	fmt.Println("Сумма введенных чисел:", sum)
}

package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Введите количество элементов массива:")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Массив в обратном порядке:")
	for i := n - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
}

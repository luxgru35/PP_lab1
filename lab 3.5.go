package main

import "fmt"

func main() {
	// Создаем массив и срез из него
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]

	// Добавляем элементы в срез
	slice = append(slice, 6, 7)
	fmt.Println("Slice после добавления:", slice)

	// Удаляем элемент из среза (например, элемент с индексом 2)
	indexToRemove := 2
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	fmt.Println("Slice после удаления:", slice)
}

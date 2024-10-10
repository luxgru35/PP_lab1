package main

import (
	"fmt"
	"unicode/utf8"
)

// Функция для определения, является ли число четным или нечетным
func evenOrOdd(num int) string {
	if num%2 == 0 {
		return "чётное"
	}
	return "нечётное"
}

// Функция для определения, является ли число положительным, отрицательным или нулем
func checkSign(num int) string {
	if num > 0 {
		return "положительно"
	} else if num < 0 {
		return "отрицательное"
	}
	return "Zero"
}

// Программа, которая выводит все числа от 1 до 10 с помощью цикла for
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// Функция для определения длины строки
func stringLength(s string) int {
	return utf8.RuneCountInString(s)
}

// Структура Rectangle
type Rectangle struct {
	width, height float64
}

// Метод для вычисления площади прямоугольника
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Функция для нахождения среднего значения двух целых чисел
func average(a, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	// Ввод числа пользователем для проверки четности
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	fmt.Printf("число: %s\n", evenOrOdd(num))

	// Проверка знака числа
	fmt.Printf("число: %s\n", checkSign(num))

	// Вывод чисел от 1 до 10
	printNumbers()

	// Проверка длины строки
	var inputString string
	fmt.Print("введите строку: ")
	fmt.Scan(&inputString)
	fmt.Printf("длина строки %d\n", stringLength(inputString))

	// Вычисление площади прямоугольника
	rect := Rectangle{width: 5.0, height: 4.0}
	fmt.Printf("площадь прямоугольника: %.2f\n", rect.area())

	// Нахождение среднего значения двух чисел
	fmt.Printf("среднее: %.2f\n", average(10, 20))
}

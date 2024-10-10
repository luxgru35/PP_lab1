package main

import (
	"fmt"
	"time"
)

// Функция для вычисления суммы и разности двух чисел с плавающей запятой
func sumAndDiff(a, b float64) (float64, float64) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	// 1. Вывод текущего времени и даты
	currentTime := time.Now()
	fmt.Println("Текущее время и дата:", currentTime)

	// 2. Объявление переменных различных типов и их вывод
	age := 25                 // int
	temperature := 36.6       // float64
	name := "А я думала сова" // string
	isStudent := true         // bool

	fmt.Println("Переменные различных типов:")
	fmt.Println("Возраст (int):", age)
	fmt.Println("Температура (float64):", temperature)
	fmt.Println("Имя (string):", name)
	fmt.Println("Студент (bool):", isStudent)

	// 3. Выполнение арифметических операций с двумя целыми числами
	a, b := 10, 3
	fmt.Println("Арифметические операции с числами", a, "и", b)
	fmt.Println("Сумма:", a+b)
	fmt.Println("Разность:", a-b)
	fmt.Println("Произведение:", a*b)
	fmt.Println("Частное:", a/b)
	fmt.Println("Остаток от деления:", a%b)

	// 4. Вызов функции для вычисления суммы и разности двух чисел с плавающей запятой
	num1, num2 := 5.5, 2.3
	sum, diff := sumAndDiff(num1, num2)
	fmt.Println("Сумма и разность чисел", num1, "и", num2)
	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", diff)

	// 5. Вычисление среднего значения трех чисел
	x, y, z := 4.5, 3.8, 5.2
	average := (x + y + z) / 3
	fmt.Println("Среднее значение чисел", x, ",", y, "и", z, "равно", average)
}

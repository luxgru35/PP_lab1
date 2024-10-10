package main

import (
	"fmt"
	"math"
)

// Функция для итерации метода Якоби
func jacobiMethod(A [][]float64, b []float64, x []float64, iterations int, tolerance float64) ([]float64, int) {
	n := len(b)
	xNew := make([]float64, n)
	iterCount := 0

	for iter := 0; iter < iterations; iter++ {
		iterCount = iter + 1
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if j != i {
					sum += A[i][j] * x[j]
				}
			}
			xNew[i] = (b[i] - sum) / A[i][i]
		}

		// Проверяем критерий сходимости
		converged := true
		for i := 0; i < n; i++ {
			if math.Abs(xNew[i]-x[i]) > tolerance {
				converged = false
				break
			}
		}
		if converged {
			break
		}

		// Обновляем x
		for i := 0; i < n; i++ {
			x[i] = xNew[i]
		}
	}

	return x, iterCount
}

func main() {
	// Коэффициенты матрицы A
	A := [][]float64{
		{30.3000, 0.0975, 0.0925, 0.0875},
		{0.1278, 29.4000, 0.1178, 0.1128},
		{0.1531, 0.1481, 28.5000, 0.1381},
		{0.1784, 0.1734, 0.1684, 27.6000},
	}

	// Свободные члены b
	b := []float64{80.1684, 83.5730, 86.6095, 89.2778}

	// Начальные приближения для x
	x := []float64{0, 0, 0, 0}

	// Количество итераций и допуск
	iterations := 100
	tolerance := 1e-6

	// Вызываем метод Якоби
	solution, iterCount := jacobiMethod(A, b, x, iterations, tolerance)

	// Выводим решение
	fmt.Println("Решение системы методом Якоби:")
	for i, val := range solution {
		fmt.Printf("x%d = %.6f\n", i+1, val)
	}

	// Выводим количество итераций
	fmt.Printf("Метод Якоби завершил работу за %d итераций.\n", iterCount)
}

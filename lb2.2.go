package main

import (
	"fmt"
	"math"
)

// Функция для итерации метода верхней релаксации (SOR)
func sorMethod(A [][]float64, b []float64, x []float64, omega float64, iterations int, tolerance float64) ([]float64, int) {
	n := len(b)
	iterCount := 0

	for iter := 0; iter < iterations; iter++ {
		iterCount = iter + 1
		xOld := make([]float64, n)
		copy(xOld, x)

		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if j != i {
					sum += A[i][j] * x[j]
				}
			}
			// Обновление x[i] с учетом коэффициента релаксации omega
			x[i] = (1-omega)*xOld[i] + omega*(b[i]-sum)/A[i][i]
		}

		// Проверка на сходимость
		converged := true
		for i := 0; i < n; i++ {
			if math.Abs(x[i]-xOld[i]) > tolerance {
				converged = false
				break
			}
		}
		if converged {
			break
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

	// Коэффициент релаксации
	omega := 1.25  // Обычно выбирается между 1 и 2

	// Количество итераций и допуск
	iterations := 100
	tolerance := 1e-6

	// Вызываем метод верхней релаксации
	solution, iterCount := sorMethod(A, b, x, omega, iterations, tolerance)

	// Выводим решение
	fmt.Println("Решение системы методом верхней релаксации (SOR):")
	for i, val := range solution {
		fmt.Printf("x%d = %.6f\n", i+1, val)
	}

	// Выводим количество итераций
	fmt.Printf("Метод верхней релаксации (SOR) завершил работу за %d итераций.\n", iterCount)
}

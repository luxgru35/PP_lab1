package mathutils

// Функция для вычисления факториала числа
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

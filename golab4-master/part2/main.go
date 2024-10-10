package main

import "fmt"

func main() {
	people := map[string]int{"John": 18, "Wick": 31}

	avg := mapAvg(people)

	fmt.Println(avg)
}

func mapAvg(people map[string]int) float64 {
	sum := 0

	for _, v := range people {
		sum += v
	}

	return float64(sum) / float64(len(people))
}

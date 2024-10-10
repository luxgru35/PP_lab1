package main

import "fmt"

func main() {
	people := map[string]int{"John": 18, "Wick": 31}

	delete(people, "John")

	for k, v := range people {
		fmt.Printf("name:%s age:%d\n", k, v)
	}
}

package main

import "fmt"

func main() {
	people := map[string]int{"John": 18, "Wick": 31}

	people["Ivan"] = 89

	for k, v := range people {
		fmt.Printf("name:%s age:%d\n", k, v)
	}
}

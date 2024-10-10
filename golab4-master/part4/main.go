package main

import (
	"fmt"
	"strings"
)

func main() {
	word := ""

	fmt.Scan(&word)

	fmt.Println(strings.ToUpper(word))
}

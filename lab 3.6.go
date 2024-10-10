package main

import "fmt"

func findLongestString(strings []string) string {
	longest := ""
	for _, str := range strings {
		if len(str) > len(longest) {
			longest = str
		}
	}
	return longest
}

func main() {
	strings := []string{"apple", "banana", "watermelon", "grape", "pineapple"}
	longest := findLongestString(strings)
	fmt.Printf("самая длинная строка: %s\n", longest)
}

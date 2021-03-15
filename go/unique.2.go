package main

import "fmt"

func unique(input string) bool {
	dictionary := map[rune]bool{}

	for _, char := range input {
		if _, ok := dictionary[char]; ok {
			return false
		}

		dictionary[char] = true
	}

	return true
}

func main() {
	fmt.Println(unique("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(unique("abcdefghijklmnopqrstuvwxyzz"))
	fmt.Println(unique("aa"))
	fmt.Println(unique("a"))
	fmt.Println(unique(""))
}

package main

import "fmt"

func unique(input string) bool {
	array := []rune(input)

	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				return false
			}
		}
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

package main

import (
	"fmt"
)

func urlify(input string, length int) string {
	array := []rune(input)
	space := []rune(" ")[0]
	zero := []rune("0")[0]
	two := []rune("2")[0]
	percentage := []rune("%")[0]

	for i, j := 0, 0; i < length; i, j = i+1, j+1 {
		original := length - 1 - i
		new := len(input) - 1 - j

		if array[original] == space {
			array[new] = zero
			array[new-1] = two
			array[new-2] = percentage

			j = j + 2
		} else {
			array[new] = array[original]
		}
	}

	return string(array)
}

func main() {
	fmt.Println(urlify("Mr John Smith    ", 13))
}

package main

import "fmt"

func plusOne(digits []int) []int {
	length := len(digits)
	flipped := false

	for i := 0; i < length; i++ {
		index := length - 1 - i

		if (digits[index]+1)%10 != 0 {
			digits[index]++

			flipped = false
			break
		}

		digits[index] = 0
		flipped = true
	}

	if flipped {
		digits = append([]int{1}, digits[:]...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{1, 2, 9}))
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{2, 9, 9}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}

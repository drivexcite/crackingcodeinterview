package main

import "fmt"

func containsDuplicate(nums []int) bool {
	dictionary := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := dictionary[nums[i]]; ok {
			return true
		} else {
			dictionary[nums[i]] = nums[i]
		}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

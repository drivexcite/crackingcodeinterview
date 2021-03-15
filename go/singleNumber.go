package main

import "fmt"

func singleNumber(nums []int) int {
	dictionary := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := dictionary[nums[i]]; ok {
			delete(dictionary, nums[i])
		} else {
			dictionary[nums[i]] = nums[i]
		}
	}

	for _, element := range dictionary {
		return element
	}

	return nums[0]
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

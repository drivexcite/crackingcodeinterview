// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	length := len(nums)
	found := make(map[int]int, length)

	for i := 0; i < length; i++ {
		complement := target - nums[i]

		if j, ok := found[complement]; ok {
			return []int{i, j}
		}

		found[nums[i]] = i
	}

	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

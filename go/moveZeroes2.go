/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
*/

package main

import "fmt"

func firstNonZero(nums []int, start int) int {

	for i := start; i < len(nums); i++ {
		if nums[i] != 0 {
			return i
		}
	}

	return -1
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func moveZeroes(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			j := firstNonZero(nums, i+1)

			if j < 0 {
				break
			}

			swap(nums, i, j)
		}
	}

	return nums
}

func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
	fmt.Println(moveZeroes([]int{0, 0, 1}))
	fmt.Println(moveZeroes([]int{0, 1, 0}))
	fmt.Println(moveZeroes([]int{0, 0}))
	fmt.Println(moveZeroes([]int{0}))
	fmt.Println(moveZeroes([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}))
	fmt.Println(moveZeroes([]int{1, 0, 0, 1, 0, 0, 1, 0, 0, 0}))
}

package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				result = append(result, nums2[j])
				nums2 = append(nums2[:j], nums2[j+1:]...)

				break
			}
		}
	}

	return result
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2}))
}

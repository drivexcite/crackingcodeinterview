package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	dictionary := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		if _, ok := dictionary[nums1[i]]; !ok {
			dictionary[nums1[i]] = 0b01
		}
	}

	for i := 0; i < len(nums2); i++ {
		if _, ok := dictionary[nums2[i]]; ok {
			dictionary[nums2[i]] = dictionary[nums2[i]] | 0b10
		}
	}

	result := make([]int, 0)

	for element, bitMask := range dictionary {
		fmt.Printf("%d, %#b\n", element, bitMask)

		if bitMask&0b11 == 0b11 {
			result = append(result, element)
		}
	}

	return result
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

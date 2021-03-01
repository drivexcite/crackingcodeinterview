package main

import "fmt"

func swap(array []int, i int, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}

func generatePermutations(array []int, size int, results *[][]int) {
	if size == 1 {
		permutation := make([]int, len(array))
		copy(permutation, array)

		*results = append(*results, permutation)
		return
	}

	for i := 0; i < size; i++ {
		generatePermutations(array, size-1, results)

		j := size - 1
		var pivot int

		if size%2 == 1 {
			pivot = 0
		} else {
			pivot = i
		}

		swap(array, pivot, j)
	}
}

func permute(nums []int) [][]int {
	results := [][]int{}
	generatePermutations(nums, len(nums), &results)

	return results
}

func main() {
	fmt.Printf("%v\n", permute([]int{1, 2, 3}))
}

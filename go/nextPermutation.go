package main

import "fmt"

func swap(array []int, i int, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}

func reverse(array []int, start int) []int {
	for i := 0; i < (len(array)-start)/2; i++ {
		swap(array, i+start, len(array)-1-i)
	}

	return array
}

func firstDecreasingFromRight(array []int) int {
	for i, j := len(array)-2, len(array)-1; i >= 0; i, j = i-1, j-1 {
		if array[i] < array[j] {
			return i
		}
	}

	return -1
}

func getMinumumLargerThan(array []int, pivotIndex int) int {
	min := int(^uint(0) >> 1)
	minIndex := -1

	for i := pivotIndex + 1; i < len(array); i++ {
		if array[i] <= min && array[i] > array[pivotIndex] {
			min = array[i]
			minIndex = i
		}
	}
	return minIndex
}

func nextPermutation(array []int) {
	pivotIndex := firstDecreasingFromRight(array)
	if pivotIndex < 0 {
		reverse(array, 0)
		return
	}

	nextIndex := getMinumumLargerThan(array, pivotIndex)
	if nextIndex < 0 {
		reverse(array, pivotIndex)
		return
	}

	swap(array, pivotIndex, nextIndex)
	reverse(array, pivotIndex+1)
}

func main() {
	array := []int{1, 5, 8, 4, 7, 6, 5, 3, 1}
	nextPermutation(array)

	fmt.Printf("%v", array)
}

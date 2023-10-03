package main

import (
	"fmt"
)

func quicksort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivot := nums[len(nums)/2]
	left := []int{}
	right := []int{}

	for i := 0; i < len(nums); i++ {
		if i == len(nums)/2 {
			continue
		}
		if nums[i] < pivot {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	left = quicksort(left)
	right = quicksort(right)

	return append(append(left, pivot), right...)
}

func main() {
	nums := []int{34, 2, 56, 76, 3, 4, 65, 7, 6, 77, 42}
	fmt.Println("Before sorting:", nums)

	sortedNums := quicksort(nums)
	fmt.Println("After sorting:", sortedNums)
}

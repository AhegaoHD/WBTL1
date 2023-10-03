package main

import (
	"fmt"
)

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	nums := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	target := 14

	index := binarySearch(nums, target)

	if index != -1 {
		fmt.Printf("Found %d at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found in the slice\n", target)
	}
}

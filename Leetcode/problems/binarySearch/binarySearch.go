package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	tagert := 9

	fmt.Println(search(nums, tagert))
}

func search(nums []int, target int) int {

	start := 0
	end := len(nums) - 1
	for start <= end {
		med := (end + start) / 2

		if nums[med] == target {
			return med
		} else if nums[med] > target {
			end = med - 1
		} else {
			start = med + 1
		}

	}

	return -1
}

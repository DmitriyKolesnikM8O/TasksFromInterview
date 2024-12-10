package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{1}
	fmt.Println(singleNumber(array))
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums)

	for i := 0; i < len(nums)-3; i += 3 {
		if nums[i] == nums[i+1] {
			continue
		} else {
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}

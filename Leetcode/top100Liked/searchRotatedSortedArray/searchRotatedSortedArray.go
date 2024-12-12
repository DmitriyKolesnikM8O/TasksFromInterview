package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 3))
}

func search(nums []int, target int) int {
	numsSortedCopy := make([]int, len(nums))
	copy(numsSortedCopy, nums)
	sort.Ints(numsSortedCopy)

	// fmt.Println(numsSortedCopy)
	l, r := 0, len(numsSortedCopy)-1
	pos := -1
	for l <= r {
		m := l + (r-l)/2
		if numsSortedCopy[m] == target {
			pos = m
			break
		} else if numsSortedCopy[m] < target {
			l = m + 1
			continue
		}
		r = m - 1
	}
	// fmt.Println(pos)

	if pos != -1 {
		for i := 0; i < len(nums); i++ {
			if numsSortedCopy[pos] == nums[i] {
				return i
			}
		}
	}

	return pos
}

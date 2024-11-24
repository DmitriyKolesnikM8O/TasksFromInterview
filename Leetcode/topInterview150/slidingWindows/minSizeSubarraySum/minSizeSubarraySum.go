package main

import (
	"fmt"
)

func main() {
	arr := []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}
	target := 213
	fmt.Println(minSubArrayLen(target, arr))
}

func minSubArrayLen(target int, nums []int) int {
	out := 0

	start, sum := 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		for sum >= target {
			if out == 0 || (i-start+1) < out {
				out = i - start + 1
			}

			sum -= nums[start]
			start++
		}
	}

	return out
}

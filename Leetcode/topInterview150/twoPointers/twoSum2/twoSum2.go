package main

import (
	"fmt"
)

func main() {
	number := []int{-10, -8, -2, 1, 2, 5, 6}
	target := 0
	fmt.Println(twoSum(number, target))
}

// func twoSum(numbers []int, target int) []int {
// 	res := []int{}

// 	for i := 0; i < len(numbers)-1; i++ {
// 		for j := i + 1; j < len(numbers); j++ {
// 			if numbers[i]+numbers[j] == target {
// 				res = append(res, i+1)
// 				res = append(res, j+1)
// 			}
// 		}
// 	}
// 	return res[:2]
// }

// binary search
func twoSum(number []int, target int) []int {

	for i := 0; i < len(number); i++ {
		j := binarySearch(number, target-number[i], i)
		if j != -1 {
			return []int{i + 1, j + 1}
		}
	}

	return []int{}
}

func binarySearch(number []int, finder int, restrict int) int {
	l := 0
	r := len(number) - 1

	for l <= r {
		m := (l + r) / 2

		if finder == number[m] && m != restrict {
			return m
		}

		if finder > number[m] {
			l = m + 1
		}
		if finder < number[m] {
			r = m - 1
		}
	}

	return -1
}

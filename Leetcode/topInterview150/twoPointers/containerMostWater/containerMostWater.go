package main

import (
	"fmt"
)

func main() {
	area := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(area))
}

func maxArea(number []int) int {
	result := 0

	i := 0
	j := len(number) - 1
	for i <= j {
		if number[i] < number[j] {
			if (number[i] * (j - i)) > result {
				result = number[i] * (j - i)
			}
		} else {
			if (number[j] * (j - i)) > result {
				result = number[j] * (j - i)
			}
		}
		if number[i] < number[j] {
			i++
		} else {
			j--
		}
	}

	return result
}

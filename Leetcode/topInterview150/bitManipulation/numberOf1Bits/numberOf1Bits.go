package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(hammingWeight(2147483645))
}

func hammingWeight(n int) int {
	stringNum := strconv.FormatInt(int64(n), 2)
	count1 := 0

	for _, i := range stringNum {
		if string(i) == "1" {
			count1++
		}
	}

	return count1
}

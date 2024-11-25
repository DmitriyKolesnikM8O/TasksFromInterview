package main

import (
	"fmt"
	"sort"
)

func main() {
	interval := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	fmt.Println(merge(interval))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals [][]int) [][]int {
	start, end := 0, 1

	result := make([][]int, 0)

	sort.Slice(intervals, func(a, b int) bool {
		return (intervals[a][0] < intervals[b][0]) || ((intervals[a][0] == intervals[b][0]) && (intervals[a][1] < intervals[b][1]))
	})

	for _, current := range intervals {
		if len(result) == 0 || (result[len(result)-1][end] < current[start]) {
			result = append(result, current)
		} else {
			result[len(result)-1][end] = max(result[len(result)-1][end], current[end])
		}
	}

	return result

}

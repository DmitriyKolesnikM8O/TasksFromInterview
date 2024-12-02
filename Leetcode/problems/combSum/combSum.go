package main

import "fmt"

func main() {
	comb := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(comb)
}

func permutation(index int, target int, candidates []int, current []int, result *[][]int) {
	if target == 0 {
		comb := make([]int, len(current))
		copy(comb, current)
		*result = append(*result, comb)
		return
	}

	for i := index; i < len(candidates); i++ {
		if candidates[i] <= target {
			current = append(current, candidates[i])

			permutation(i, target-candidates[i], candidates, current, result)

			current = current[:len(current)-1]
		}
	}

}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	current := []int{}
	permutation(0, target, candidates, current, &result)

	return result
}

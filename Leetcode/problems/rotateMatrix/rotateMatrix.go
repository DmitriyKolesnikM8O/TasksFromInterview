package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	rotate(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j < i {
				temp := matrix[i][j]
				matrix[i][j] = matrix[j][i]
				matrix[j][i] = temp
			}

		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])/2; j++ {

			temp := matrix[i][j]
			matrix[i][j] = matrix[i][len(matrix)-j-1]
			matrix[i][len(matrix)-j-1] = temp

		}
	}

}

/*
[0][0] -> [0][2]
[0][1] -> [1][2]
[0][2] -> [2][2]
[1][0] -> [0][1]
[1][1] -> [1][1]
[1][2] -> [2][1]
[2][0] -> [0][0]
[2][1] -> [1][0]
[2][2] -> [2][0]
*/

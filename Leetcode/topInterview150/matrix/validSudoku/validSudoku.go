package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {

	//rows
	for i := 0; i < len(board); i++ {
		mapCount := make(map[byte]int)
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}

			if _, ok := mapCount[board[i][j]]; ok {
				return false
			}
			mapCount[board[i][j]] = mapCount[board[i][j]] + 1
		}
	}

	//columns
	for j := 0; j < 9; j++ {
		mapCount := make(map[byte]int)
		for i := 0; i < 9; i++ {
			if board[i][j] == '.' {
				continue
			}
			_, ok := mapCount[board[i][j]]
			if ok { // already exists
				return false
			}
			mapCount[board[i][j]] = mapCount[board[i][j]] + 1
		}
	}

	//3x3
	in := 0
	jn := 0
	i := 0
	j := 0

	result := true
	for in < 9 {
		for jn < 9 {
			result = check(board, i+in, j+jn, in+3, jn+3)
			if !result {
				return false
			}
			jn += 3
		}
		jn = 0
		in += 3
	}

	return true
}

func check(board [][]byte, i int, j int, in int, jn int) bool {
	mapCount := make(map[byte]int)

	for ; i < in; i++ {
		for ; j < jn; j++ {
			if board[i][j] == '.' {
				continue
			}

			if _, ok := mapCount[board[i][j]]; ok {
				return false
			}

			mapCount[board[i][j]] = mapCount[board[i][j]] + 1
		}
		j -= 3
	}

	return true
}

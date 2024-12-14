package main

import "fmt"

func main() {
	grid := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++

				bsd(grid, i, j)
			}
		}
	}

	return count
}

func bsd(grid [][]byte, s int, l int) {
	if s < 0 || s >= len(grid) || l < 0 || l >= len(grid[s]) || grid[s][l] != '1' {
		return
	}

	grid[s][l] = '2'
	bsd(grid, s+1, l)
	bsd(grid, s-1, l)
	bsd(grid, s, l+1)
	bsd(grid, s, l-1)
}

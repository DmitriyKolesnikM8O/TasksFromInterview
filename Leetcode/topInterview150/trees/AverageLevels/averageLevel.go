package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   20,
			Right: nil,
			Left:  nil,
		},
	}

	fmt.Println(averageOfLevels(tree))

}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	CurrentLevel := []*TreeNode{root}
	AverageLevels := []float64{}

	for len(CurrentLevel) != 0 {
		NextLevel := []*TreeNode{}
		Average := 0.0

		for _, node := range CurrentLevel {
			if node != nil {
				Average += float64(node.Val)
			}

			if node.Left != nil {
				NextLevel = append(NextLevel, node.Left)
			}

			if node.Right != nil {
				NextLevel = append(NextLevel, node.Right)
			}
		}

		curAverage := Average / float64(len(CurrentLevel))

		AverageLevels = append(AverageLevels, curAverage)

		CurrentLevel = NextLevel
	}

	return AverageLevels
}

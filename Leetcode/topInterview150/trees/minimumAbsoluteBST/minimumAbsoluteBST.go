package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// tree := &TreeNode{
	// 	Val: 4,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val:   1,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: &TreeNode{
	// 			Val:   3,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val:   6,
	// 		Right: nil,
	// 		Left:  nil,
	// 	},
	// }
	// tree := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val:   0,
	// 		Left:  nil,
	// 		Right: nil,
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 48,
	// 		Right: &TreeNode{
	// 			Val:   12,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Left: &TreeNode{
	// 			Val:   49,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 	},
	// }

	tree := &TreeNode{
		Val: 543,
		Left: &TreeNode{
			Val:  384,
			Left: nil,
			Right: &TreeNode{
				Val:   445,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 652,
			Right: &TreeNode{
				Val:   699,
				Left:  nil,
				Right: nil,
			},
			Left: nil,
		},
	}
	fmt.Println(getMinimumDifference(tree))
	fmt.Println(getMinimumDifference2(tree))

}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root != nil && root.Left == nil && root.Right == nil {
		return root.Val
	}

	nodeList := []int{}

	currentNode := []*TreeNode{root}

	for len(currentNode) != 0 {
		nextNode := []*TreeNode{}

		for _, node := range currentNode {
			if node != nil {
				nodeList = append(nodeList, node.Val)
			}

			if node.Left != nil {
				nextNode = append(nextNode, node.Left)
			}

			if node.Right != nil {
				nextNode = append(nextNode, node.Right)
			}
		}
		currentNode = nextNode
	}
	result := int(math.Abs(float64(nodeList[0]) - float64(nodeList[1])))
	for i := range nodeList {
		for j := range nodeList {

			if int(math.Abs(float64(nodeList[i])-float64(nodeList[j]))) < result && i != j {
				result = int(math.Abs(float64(nodeList[i]) - float64(nodeList[j])))
			}
		}
	}

	return result
}

func getMinimumDifference2(root *TreeNode) int {
	vals := []int{}

	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		vals = append(vals, root.Val)
		traverse(root.Right)
	}
	traverse(root)

	minDiff := math.MaxInt64
	for i := 1; i < len(vals); i++ {
		minDiff = min(minDiff, vals[i]-vals[i-1])
	}

	return minDiff
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

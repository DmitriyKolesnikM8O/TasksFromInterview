package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
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
	}

	fmt.Println(isBalanced(tree))
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(height(root.Left), height(root.Right))
}

func abs(one int, two int) int {
	if one > two {
		return one - two
	}
	return two - one
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := height(root.Left)
	rh := height(root.Right)

	if abs(lh, rh) < 2 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}

	return false
}

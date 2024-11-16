package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	targetSumm := 22
	fmt.Println(hasPathSum(tree, targetSumm))

}

func CheckerPaths(root *TreeNode, result *bool, targetSum int) {
	if *result {
		return
	}

	if root.Left != nil {
		root.Left.Val = root.Val + root.Left.Val
		CheckerPaths(root.Left, result, targetSum)
	}

	if root.Right != nil {
		root.Right.Val = root.Val + root.Right.Val
		CheckerPaths(root.Right, result, targetSum)
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			*result = true
		}
	}
	return
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	result := false
	CheckerPaths(root, &result, targetSum)
	return result
}

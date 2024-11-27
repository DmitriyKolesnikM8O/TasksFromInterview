package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	fmt.Println(postorderTraversal(tree))
}

func postorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	result = append(result, root.Val)
	result = append(result, postorder(root.Right)...)
	result = append(result, postorder(root.Left)...)

	return result
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	result = append(result, postorder(root)...)

	// fmt.Println(result)

	resultReverse := []int{}
	for i := len(result) - 1; i >= 0; i-- {
		resultReverse = append(resultReverse, result[i])
	}

	return resultReverse
}

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// tree := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val:   4,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: &TreeNode{
	// 			Val:   5,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 3,
	// 		Left: &TreeNode{
	// 			Val:   6,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: nil,
	// 	},
	// }

	tree := &TreeNode{}
	fmt.Println(countNodes(tree))

}

func Counter(root *TreeNode, result *int) {
	if root != nil {
		*result++
	}

	if root.Left != nil {
		Counter(root.Left, result)
	}
	if root.Right != nil {
		Counter(root.Right, result)
	}

	return
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	Counter(root, &result)

	return result
}

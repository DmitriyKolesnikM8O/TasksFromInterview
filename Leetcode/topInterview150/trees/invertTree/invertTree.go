package main

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
	// 		Val: 7,
	// 		Left: &TreeNode{
	// 			Val:   6,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: &TreeNode{
	// 			Val:   9,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 	},
	// }

	tree := &TreeNode{}
	_ = invertTree(tree)

}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{}
	}

	NewTree := &TreeNode{
		Val:   root.Val,
		Left:  nil,
		Right: nil,
	}

	if root.Right != nil {
		NewTree.Left = invertTree(root.Right)
	}
	if root.Left != nil {
		NewTree.Right = invertTree(root.Left)
	}

	return NewTree
}

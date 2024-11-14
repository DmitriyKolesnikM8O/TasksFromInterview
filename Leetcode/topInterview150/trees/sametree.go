package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	p := TreeNode{}

	q := TreeNode{}

	fmt.Println(isSameTree(&p, &q))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	result := true
	if p.Val != q.Val {
		result = false
	}

	if p.Left != nil && q.Left != nil && result != false {
		result = isSameTree(p.Left, q.Left)
	}
	if (p.Left != nil && q.Left == nil) || (p.Left == nil && q.Left != nil) {
		result = false
	}

	if p.Right != nil && q.Right != nil && result != false {
		result = isSameTree(p.Right, q.Right)
	}
	if (p.Right != nil && q.Right == nil) || (p.Right == nil && q.Right != nil) {
		result = false
	}

	return result
}

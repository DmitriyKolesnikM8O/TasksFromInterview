package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(isSymmetric(tree))

}

func InOrderTraversalLeft(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	var result []int
	result = append(result, InOrderTraversalLeft(node.Left)...)
	result = append(result, node.Val)
	result = append(result, InOrderTraversalLeft(node.Right)...)

	return result
}

func InOrderTraversalRight(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	var result []int
	result = append(result, InOrderTraversalRight(node.Right)...)
	result = append(result, node.Val)
	result = append(result, InOrderTraversalRight(node.Left)...)

	return result
}

func isSymmetric(root *TreeNode) bool {
	leftNodes := root.Left
	leftNodesList := []int{}
	for leftNodes != nil {
		leftNodesList = append(leftNodesList, leftNodes.Val)
		leftNodes = leftNodes.Left
	}

	rightNodes := root.Right
	rightNodesList := []int{}
	for rightNodes != nil {
		rightNodesList = append(rightNodesList, rightNodes.Val)
		rightNodes = rightNodes.Right
	}

	if len(rightNodesList) != len(leftNodesList) {
		return false
	}

	if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		return false
	}
	leftValues := InOrderTraversalLeft(root.Left)
	rightValues := InOrderTraversalRight(root.Right)

	fmt.Println(leftValues)
	fmt.Println(rightValues)

	if len(leftValues) != len(rightValues) {
		return false
	}

	for i := 0; i < len(leftValues); i++ {
		if leftValues[i] != rightValues[i] {
			return false
		}
	}

	return true
}

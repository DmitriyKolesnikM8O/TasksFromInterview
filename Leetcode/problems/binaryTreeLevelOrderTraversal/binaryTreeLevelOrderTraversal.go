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

	fmt.Println(levelOrder(tree))
}

func iteration(root *TreeNode, myMap map[int][]int, level int) {
	if root == nil {
		return
	}

	myMap[level] = append(myMap[level], root.Val)
	iteration(root.Left, myMap, level+1)
	iteration(root.Right, myMap, level+1)

}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	mapForMe := make(map[int][]int)
	iteration(root, mapForMe, 0)
	result := make([][]int, len(mapForMe))
	for k, v := range mapForMe {
		result[k] = v
	}

	return result
}

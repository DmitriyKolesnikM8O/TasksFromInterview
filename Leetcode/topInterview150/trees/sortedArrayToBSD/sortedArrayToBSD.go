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
	printTree(sortedArrayToBST([]int{-10, -3, 0, 5, 9}), 0)

}

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTOneIteration(0, len(nums)-1, nums)
}

func sortedArrayToBSTOneIteration(start, end int, nums []int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (end + start) / 2

	var root TreeNode
	root.Val = nums[mid]
	root.Left = sortedArrayToBSTOneIteration(start, mid-1, nums)
	root.Right = sortedArrayToBSTOneIteration(mid+1, end, nums)
	return &root
}

func printTree(root *TreeNode, tab int) {
	if root != nil {
		if root.Left != nil {
			printTree(root.Left, tab+1)
		}
		if root.Right != nil {
			printTree(root.Right, tab+1)
		}

		if tab > 0 {
			fmt.Printf("\t")
			tab -= 1
		}

		fmt.Println(root.Val)
	}
}

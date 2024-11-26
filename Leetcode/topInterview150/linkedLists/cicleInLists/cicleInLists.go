package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(value int) *ListNode {
	return &ListNode{Val: value, Next: nil}
}

func main() {
	node1 := NewNode(3)
	node2 := NewNode(2)
	node3 := NewNode(0)
	node4 := NewNode(-4)

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	fmt.Println(hasCycle(node1))
}

func hasCycle(head *ListNode) bool {
	visitednodes := make(map[*ListNode]bool)

	current := head

	for current != nil {
		if visitednodes[current] {
			return true
		}
		visitednodes[current] = true
		current = current.Next
	}

	return false
}

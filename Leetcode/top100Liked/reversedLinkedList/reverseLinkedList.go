package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	myList := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(reverseList(myList))
}

func reverseList(head *ListNode) *ListNode {

	current := head // curr starts as head
	// next := &ListNode{}
	var prev *ListNode // prev starts as nil
	for current != nil {
		next := current.Next //save next node
		current.Next = prev  //reverse
		prev = current       //move prev to next node

		current = next // mode curr to next node
	}

	return prev
}

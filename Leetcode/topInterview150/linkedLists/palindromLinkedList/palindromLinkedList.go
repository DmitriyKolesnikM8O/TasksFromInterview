package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	_ = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	fmt.Println(isPalindrome(&ListNode{}))

}

func isPalindrome(head *ListNode) bool {
	list := []int{}

	current := head
	for current != nil {
		list = append(list, current.Val)
		current = current.Next
	}

	i := 0
	j := len(list) - 1
	for i <= j {
		if list[i] != list[j] {
			return false
		}
		i++
		j--
	}

	return true

}

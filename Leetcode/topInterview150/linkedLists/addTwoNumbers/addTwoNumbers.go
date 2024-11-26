package main

import (
	"fmt"
	"math/big"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l3 := addTwoNumbers(l1, l2)

	for l3.Next != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
	fmt.Print(l3.Val)
}

func intToList(l1 *ListNode) []int {
	l1List := []int{}
	for l1.Next != nil {
		l1List = append(l1List, l1.Val)
		l1 = l1.Next
	}
	l1List = append(l1List, l1.Val)

	return l1List
}

func addInList(l *ListNode, number *big.Int) *ListNode {
	if number.Cmp(big.NewInt(0)) == 0 {
		return &ListNode{}
	}

	mod := big.NewInt(0)
	mod.Mod(number, big.NewInt(10))
	l.Val = int(mod.Int64())

	number.Quo(number, big.NewInt(10))

	l.Next = &ListNode{}
	return addInList(l.Next, number)
}

func removeLastNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // Список пуст или содержит только один узел
		return nil
	}

	current := head
	for current.Next.Next != nil { // Идём до предпоследнего узла
		current = current.Next
	}
	current.Next = nil // Удаляем последний узел, обнуляя ссылку
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1List := intToList(l1)
	l2List := intToList(l2)

	l1Digit := big.NewInt(0)
	l2Digit := big.NewInt(0)

	for i := len(l1List) - 1; i >= 0; i-- {
		l1Digit = l1Digit.Mul(l1Digit, big.NewInt(10)).Add(l1Digit, big.NewInt(int64(l1List[i])))
	}
	for i := len(l2List) - 1; i >= 0; i-- {
		l2Digit = l2Digit.Mul(l2Digit, big.NewInt(10)).Add(l2Digit, big.NewInt(int64(l2List[i])))
	}

	result := l1Digit.Add(l1Digit, l2Digit)

	resultList := &ListNode{}

	addInList(resultList, result)

	final := removeLastNode(resultList)

	if final == nil {
		return &ListNode{
			Val: 0,
		}
	}
	return final

}

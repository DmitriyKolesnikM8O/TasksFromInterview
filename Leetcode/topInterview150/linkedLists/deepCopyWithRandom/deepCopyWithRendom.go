package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	myMap := make(map[*Node]*Node)

	currentNode := head
	for currentNode != nil {
		myMap[currentNode] = &Node{
			Val: currentNode.Val,
		}
		currentNode = currentNode.Next
	}

	currentNode = head
	for currentNode != nil {
		myMap[currentNode].Next = myMap[currentNode.Next]
		myMap[currentNode].Random = myMap[currentNode.Random]

		currentNode = currentNode.Next
	}

	return myMap[head]
}

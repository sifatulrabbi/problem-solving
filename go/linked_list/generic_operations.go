package linked_list

import "fmt"

// Insert 60 after 5 in the linked list.
func Insertion() {
	nodeList := getNumbersLinkedList()
	fmt.Println("Inserting 60 after 5 in the linked list.")
	fmt.Println("Before:", convertIntoArr(nodeList))

	node := &nodeList
	for {
		if node.Val == 5 {
			break
		}
		node = node.Next
	}
	next := node.Next
	newNode := &DoubleLinkedList[int]{
		Val:  60,
		Next: next,
		Prev: node,
	}
	next.Prev = newNode
	node.Next = newNode

	fmt.Println("After:", convertIntoArr(nodeList))
}

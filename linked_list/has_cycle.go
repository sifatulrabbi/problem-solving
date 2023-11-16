package linked_list

import "fmt"

type listNode struct {
	val  int
	next *listNode
}

func RunHasCycle() {
	head := createCycledLinkedList([]int{3, 2, 0, -4}, 1)
	fmt.Println("has cycle -", hasCycle(&head))

	head = createCycledLinkedList([]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}, -1)
	fmt.Println("has cycle -", hasCycle(&head))

	head = createCycledLinkedList([]int{1}, -1)
	fmt.Println("no cycle -", hasCycle(&head))
}

func hasCycle(head *listNode) bool {
	tortoise := head
	hare := head
	for hare.next != nil && hare.next.next != nil && hare.next.next.next != nil {
		tortoise = tortoise.next
		hare = hare.next.next.next
		if tortoise == hare {
			return true
		}
	}
	return false
}

func createCycledLinkedList(arr []int, pos int) listNode {
	head := &listNode{
		val:  arr[0],
		next: nil,
	}
	current := head
	var ref *listNode = nil

	if pos == 0 {
		ref = head
	}

	for i := 1; i < len(arr); i++ {
		v := arr[i]
		newNode := &listNode{
			val:  v,
			next: nil,
		}
		current.next = newNode
		current = newNode
		if pos == i {
			ref = newNode
		}
	}
	if ref != nil {
		current.next = ref
	}
	return *head
}

package linkedlist

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func createtLinkedList(length int) *ListNode2 {
	var head *ListNode2
	var curNode *ListNode2
	for i := 0; i < length; i++ {
		newNode := ListNode2{i, nil}
		if head == nil {
			head = &newNode
			curNode = head
		} else {
			curNode.Next = &newNode
			curNode = curNode.Next
		}
	}
	return head
}

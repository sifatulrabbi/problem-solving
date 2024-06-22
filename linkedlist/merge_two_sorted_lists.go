package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		head      *ListNode = nil
		curNode   *ListNode = nil
		nextList1 *ListNode = list1
		nextList2 *ListNode = list2
	)

	for nextList1 != nil || nextList2 != nil {
		if nextList2 != nil && nextList1 == nil {
			curNode = &ListNode{nextList2.Val, nil}
		} else if nextList1 != nil && nextList2 == nil {
			curNode = &ListNode{nextList1.Val, nil}
		} else if nextList1.Val < nextList2.Val {
			curNode = &ListNode{nextList2.Val, nil}
		} else {
			curNode = &ListNode{nextList1.Val, nil}
		}

		if nextList2 != nil {
			nextList2 = nextList2.Next
		}
		if nextList1 != nil {
			nextList1 = nextList1.Next
		}

		curNode.Next = nil
		if head == nil {
			head = curNode
		} else {
			head.Next = curNode
			curNode = head.Next
		}
	}

	return head
}

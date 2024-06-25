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
		var selNode ListNode
		if nextList1 != nil && nextList2 != nil {
			if nextList1.Val < nextList2.Val {
				selNode = ListNode{nextList1.Val, nil}
				nextList1 = nextList1.Next
			} else {
				selNode = ListNode{nextList2.Val, nil}
				nextList2 = nextList2.Next
			}
		} else if nextList1 != nil && nextList2 == nil {
			selNode = ListNode{nextList1.Val, nil}
			nextList1 = nextList1.Next
		} else if nextList1 == nil && nextList2 != nil {
			selNode = ListNode{nextList2.Val, nil}
			nextList2 = nextList2.Next
		}

		if head == nil {
			head = &selNode
			curNode = head
		} else {
			curNode.Next = &selNode
			curNode = curNode.Next
		}
	}

	return head
}

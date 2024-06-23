package linkedlist

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	exmp1List1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	exmp1List2 := &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}
	head := mergeTwoLists(exmp1List1, exmp1List2)
	fmt.Println(head)

	headItem := &ListNode{1, nil}
	var curItem *ListNode
	item1 := &ListNode{2, &ListNode{3, &ListNode{4, nil}}}
	fmt.Println(headItem, curItem, item1)

	curItem = headItem
	curItem.Next = item1
	fmt.Printf("%+v | %+v | %+v\n", headItem, curItem, item1)

	curItem = curItem.Next
	curItem.Next = item1.Next
	fmt.Printf("%+v | %+v | %+v\n", headItem, curItem, item1)
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		return
	}
	node.Val = node.Next.Val
	if node.Next.Next == nil {
		node.Next = nil
	}
	deleteNode(node.Next)
}

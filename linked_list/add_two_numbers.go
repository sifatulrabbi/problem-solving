package linked_list

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func RunAddTwoNumbers() {
	// testcase-1
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	res := addTwoNumbers(l1, l2)
	fmt.Println(convertListNodeToNum(res))

	// testcase-2
	l1 = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9, Next: nil}}}
	l2 = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9, Next: nil}}}}
	res = addTwoNumbers(l1, l2)
	fmt.Println(convertListNodeToNum(res))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := convertListNodeToNum(l1)
	num2 := convertListNodeToNum(l2)
	resultNum := new(big.Int)
	resultNum = resultNum.Add(num1, num2)
	sumStr := resultNum.Text(10)
	l := len(sumStr)

	lastNum, _ := strconv.ParseInt(string(sumStr[l-1]), 10, 32)
	head := &ListNode{
		Val:  int(lastNum),
		Next: nil,
	}
	current := head
	for i := l - 2; i >= 0; i-- {
		n, _ := strconv.ParseInt(string(sumStr[i]), 10, 32)
		nextNode := &ListNode{
			Val:  int(n),
			Next: nil,
		}
		current.Next = nextNode
		current = nextNode
	}
	return head
}

func convertListNodeToNum(l *ListNode) *big.Int {
	n := new(big.Int)
	val := ""
	current := l
	for current != nil {
		val += fmt.Sprintf("%d", current.Val)
		current = current.Next
	}
	n, ok := n.SetString(val, 10)
	if !ok {
		panic(errors.New("the string is not a valid number"))
	}
	return n
}

// utilities

func createListNodeFromNumbers(nums []int) *ListNode {
	head := &ListNode{Val: nums[0], Next: nil}
	current := head
	for i := 1; i < len(nums); i++ {
		nextNode := &ListNode{Val: nums[i], Next: nil}
		current.Next = nextNode
		current = nextNode
	}
	return head
}

func createArrFromListNode(head *ListNode) []int {
	arr := []int{}
	current := head
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}
	return arr
}

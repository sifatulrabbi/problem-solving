package linked_list

import (
	"fmt"
	"log"
)

func RunTestQueue() {
	arr := []int{31, 24, 32, 16, 23, 15, 35, 27, 32}
	q := BuildQueueFromArr[int](arr)

	fmt.Println(arr)
	fmt.Println("Getting the first value from the queue which should be 31")
	if v, exists := q.Dequeue(); !exists {
		log.Fatalf("The queue is empty")
	} else {
		fmt.Printf("value found: %d. Which matches: %v\n", v, v == 31)
	}

	fmt.Println("Getting the next value from the queue which should be 24")
	if v, exists := q.Dequeue(); !exists {
		log.Fatalf("The queue is empty")
	} else {
		fmt.Printf("value found: %d. Which matches: %v\n", v, v == 24)
	}
	fmt.Println(ConvertQueueIntoArray[int](q))
}

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

type Queue[T comparable] struct {
	Length int
	Head   *Node[T]
	Tail   *Node[T]
}

func (q *Queue[T]) Enqueue(item T) {
	n := Node[T]{
		Val:  item,
		Next: nil,
	}

	if q.Length == 0 {
		// when the queue length is 0 its empty.
		q.Head = &n
		q.Tail = &n
		q.Length++
		return
	}

	q.Tail.Next = &n
	q.Tail = &n
	q.Length++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.Head == nil {
		var n T
		return n, false
	}

	head := q.Head
	q.Head = head.Next
	q.Length--
	return head.Val, false
}

func (q *Queue[T]) Peek() (T, bool) {
	var val T
	if q.Head != nil {
		return q.Head.Val, true
	}
	return val, false
}

func ConvertQueueIntoArray[T comparable](q Queue[T]) []T {
	arr := []T{}
	for {
		if val, exists := q.Dequeue(); !exists {
			break
		} else {
			arr = append(arr, val)
		}
	}
	return arr
}

func BuildQueueFromArr[T comparable](arr []T) Queue[T] {
	q := Queue[T]{}
	for _, v := range arr {
		q.Enqueue(v)
	}
	return q
}

package linked_list

type DoubleLinkedList[T comparable] struct {
	Val  T
	Next *DoubleLinkedList[T]
	Prev *DoubleLinkedList[T]
}

func getNumbersLinkedList() DoubleLinkedList[int] {
	head := &DoubleLinkedList[int]{
		Val:  1,
		Prev: nil,
		Next: nil,
	}
	current := head
	for i := 2; i < 20; i++ {
		newNode := &DoubleLinkedList[int]{
			Val:  i,
			Prev: current,
			Next: nil,
		}
		current.Next = newNode
		current = newNode
	}
	return *head
}

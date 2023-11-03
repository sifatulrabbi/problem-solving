package linked_list

import "fmt"

func ConvertLinkedListToArray() {
	node := getNumbersLinkedList()
	arr := []int{}
	for {
		arr = append(arr, node.Val)
		if node.Next == nil {
			break
		} else {
			node = *node.Next
		}
	}
	fmt.Println(arr)
}

func convertIntoArr(nodeList DoubleLinkedList[int]) []int {
	node := nodeList
	arr := []int{}
	for {
		arr = append(arr, node.Val)
		if node.Next == nil {
			break
		} else {
			node = *node.Next
		}
	}
	return arr
}

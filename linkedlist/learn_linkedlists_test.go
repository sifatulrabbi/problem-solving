package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLinkedList(t *testing.T) {
	head := createtLinkedList(5)
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	fmt.Println(values)
	assert.True(t, true, len(values) == 5)
}

package turingproblems

import (
	"fmt"
	"testing"
)

func TestLargestNumInArr(t *testing.T) {
	fmt.Println("Testing with:", 7)
	result := largestNumInArr(7)
	fmt.Println("Result:", result)
	if result != 3 {
		t.Errorf("Result should be 3, got: %d\n", result)
	}
}

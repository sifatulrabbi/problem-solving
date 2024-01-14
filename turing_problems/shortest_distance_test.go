package turingproblems

import (
	"fmt"
	"testing"
)

func TestShortestDistance(t *testing.T) {
	fmt.Println("testing with [[1,1], [6,2], [1,5], [3,1]]")
	x, y := 5, 1
	arr := [][]int{{1, 1}, {6, 2}, {1, 5}, {3, 1}}
	distance := shortestDistance(x, y, arr)
	if distance != 3 {
		t.Errorf("Result should be 3, got: %d\n", distance)
	}

	x, y = 3, 4
	arr = [][]int{{2, 3}}
	distance = shortestDistance(x, y, arr)
	if distance != -1 {
		t.Errorf("Result should be -1, got: %d\n", distance)
	}
}

package others

import (
	"testing"
)

func TestCarFleet(t *testing.T) {
	testCases := []struct {
		target         int
		position       []int
		speed          []int
		expectedFleets int
	}{
		{10, []int{4, 1, 0, 7}, []int{2, 2, 1, 1}, 3},
		{10, []int{1, 4}, []int{3, 2}, 1},
	}
	for _, tc := range testCases {
		if fleets := carFleet(tc.target, tc.position, tc.speed); fleets != tc.expectedFleets {
			t.Fatalf("target=%d position=%v speed=%v -> expected=%d got=%d\n",
				tc.target, tc.position, tc.speed, tc.expectedFleets, fleets)
		}
	}
}

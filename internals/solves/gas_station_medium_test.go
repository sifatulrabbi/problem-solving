package solves

import "testing"

func TestGasStationMedium(t *testing.T) {
	t.Skip()
	idx := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	t.Log(idx)
	idx = canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})
	t.Log(idx)
	idx = canCompleteCircuit([]int{5, 8, 2, 8}, []int{6, 5, 6, 6})
	t.Log(idx)
	idx = canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2})
	t.Log(idx)
}

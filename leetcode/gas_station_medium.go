package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) < 1 {
		return -1
	}

	startIdx := -1
	for i := 0; i < len(gas); i++ {
		if gas[i] < cost[i] {
			continue
		}
		tank := gas[i]
		curPos, nextPos := i, i+1
		for nextPos != i {
			if nextPos > len(gas)-1 {
				nextPos = nextPos - len(gas)
			}
			tank += gas[nextPos] - cost[curPos]
			if tank < cost[nextPos] {
				break
			}
			curPos = nextPos
			nextPos++
		}
		if nextPos == i {
			startIdx = i
			break
		}
	}
	return startIdx
}

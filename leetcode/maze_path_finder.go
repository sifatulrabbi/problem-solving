package leetcode

var SampleMaze = []string{
	"X", "X", "X", "E", "X",
	"X", " ", " ", " ", "X",
	"X", " ", "X", "X", "X",
	"X", "S", "X", "X", "X",
}

func walk(maze []string, curr, start, end, wall, path string) bool {
	// base cases
	// 1. if off the map
	if curr == start {
		return false
	}
	if curr == end {
		return false
	}
	if curr == wall {
		return false
	}
	if curr == path {
		return true
	}
	// 2. if hit a wall
	// 3. if at end
	// 4. if already seen
	// 5. if found path
	return false
}

func MazeSolver() {
}

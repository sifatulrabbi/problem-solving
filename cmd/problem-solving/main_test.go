package main

import (
	"testing"
)

func TestMainFunc(t *testing.T) {
	// activeCus := mostActive([]string{"Omega", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Omega", "Alpha", "Beta"})
	l := longestSubarray([]int32{1, 1, 1, 3, 3, 2, 2})
	t.Log(l)
	l = longestSubarray([]int32{5, 1, 2, 3, 4, 5})
	t.Log(l)
	l = longestSubarray([]int32{3, 2, 2, 1})
	t.Log(l)
	l = longestSubarray([]int32{31, 157793605, 157793605, 157793604, 157793604, 157793604, 157793604, 157793605, 157793605, 157793605, 157793604, 157793605, 157793604, 157793605, 157793605, 157793604, 157793604, 157793604, 157793605, 157793605, 157793605, 157793605, 157793604, 157793604, 157793605, 157793604, 157793605, 157793605, 157793605, 157793604, 157793605, 157793605})
	t.Log(l)
}

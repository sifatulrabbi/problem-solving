package turingproblems

import (
	"math"
)

/*
Given your current position on a Cartesian plane, represented by the integers x and y, along with an array called "locations" where each element locations[i] = [xi, yi] signifies the presence of a point at coordinates (ai, bi), we need to find the index (starting from 0) of the valid point that is closest to your current location according to the Manhattan distance.
A point is considered valid if it shares either the same x-coordinate or the same y-coordinate as your current position. If there are multiple valid points with the same minimum Manhattan distance, we return the valid point with the smallest index. If no valid points are found, the function returns -1.

Example 1:
Input: x = 5, y = 1, locations = [[1,1], [6,2], [1,5], [3,1]]
Output: 3
Explanation: Of all points, only [3,1] are valid. Of the valid point and have the smallest Manhattan distance from your current location, so return 3.

Example 2:
Input: x = 3, y = 4, points = [[2,3]]
Output: -1
Explanation: There are no valid points.
*/
func shortestDistance(x, y int, locations [][]int) int {
	locs := [][]int{}

	for _, l := range locations {
		if l[0] == x || l[1] == y {
			locs = append(locs, l)
		}
	}
	if len(locs) < 1 {
		return -1
	}

	distance := math.Inf(1)
	point := -1
	for _, l := range locs {
		dis := math.Abs(float64(x-l[0])) + math.Abs(float64(y-l[1]))
		if dis < distance {
			distance = dis
			point = l[0]
		}
	}
	return point
}

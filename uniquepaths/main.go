package uniquepaths

/*
https://leetcode.com/problems/unique-paths/

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?
*/

func uniquePaths(m int, n int) int {
	paths := [][]int{}

	for i := 0; i < m; i++ {
		paths = append(paths, make([]int, n))
	}

	paths[0][0] = 1

	for x := 1; x < m; x++ {
		paths[x][0] = 1
	}

	for y := 1; y < n; y++ {
		paths[0][y] = 1
	}

	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			paths[x][y] = paths[x-1][y] + paths[x][y-1]
		}
	}

	return paths[m-1][n-1]
}

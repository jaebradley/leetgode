package minimumpathsum

/*
https://leetcode.com/problems/minimum-path-sum/

Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.
*/

func minPathSum(grid [][]int) int {
	length := len(grid)
	width := len(grid[0])

	sums := [][]int{}

	for i := 0; i < length; i++ {
		sums = append(sums, make([]int, width))
	}

	sums[0][0] = grid[0][0]

	for x := 0; x < length; x++ {
		for y := 0; y < width; y++ {
			currentValue := grid[x][y]

			if x > 0 && y > 0 {
				previousLeft := sums[x-1][y]
				previousUp := sums[x][y-1]

				if previousLeft > previousUp {
					sums[x][y] = previousUp + currentValue
				} else {
					sums[x][y] = previousLeft + currentValue
				}
			} else if x == 0 && y > 0 {
				sums[x][y] = sums[x][y-1] + currentValue
			} else if y == 0 && x > 0 {
				sums[x][y] = sums[x-1][y] + currentValue
			}
		}
	}

	return sums[length-1][width-1]
}

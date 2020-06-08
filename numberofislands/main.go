package numberofislands

/*
https://leetcode.com/problems/number-of-islands/

Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input:
11110
11010
11000
00000

Output: 1
Example 2:

Input:
11000
11000
00100
00011

Output: 3
*/

func numIslands(grid [][]byte) int {
	islands := 0

	if grid != nil {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == byte('1') {
					islands++
					dfs(x, y, grid)
				}
			}
		}
	}

	return islands
}

func dfs(x int, y int, grid [][]byte) {
	if grid[y][x] == byte('1') {
		grid[y][x] = byte('#')

		if x > 0 {
			dfs(x-1, y, grid)
		}

		if y > 0 {
			dfs(x, y-1, grid)
		}

		if y < len(grid)-1 {
			dfs(x, y+1, grid)
		}

		if x < len(grid[y])-1 {
			dfs(x+1, y, grid)
		}
	}
}

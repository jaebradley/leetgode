package pathsumthree

/*
https://leetcode.com/problems/path-sum-iii/

You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

Example:

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

Return 3. The paths that sum to 8 are:

1.  5 -> 3
2.  5 -> 2 -> 1
3. -3 -> 11

Approach:

* Use DFS outlined here: https://leetcode.com/problems/path-sum-iii/discuss/141424/Python-step-by-step-walk-through.-Easy-to-understand.-Two-solutions-comparison.-%3A-)
* The idea is that keep track of different path sums, and theif frequencies, as tree is traversed
* If there exist cached path sums that are equal to the current path sum minus the target sum then add their frequencies to the number of total matches
* In the example case, 10 + 5 + 3 - target (8) is 10, which indicates there's a valid path from 10 to 3 that matches the target
* When switching branches, decrement current path sum from cache count
*/

// TreeNode represents binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	paths := 0
	pathSumsCache := map[int]int{0: 1}
	helper(&paths, root, pathSumsCache, sum, 0)
	return paths
}

func helper(paths *int, node *TreeNode, pathSumsCache map[int]int, target int, pathSum int) {
	if node != nil {
		nextPathSum := pathSum + node.Val
		previousSumCount, previousSumExists := pathSumsCache[nextPathSum-target]
		if previousSumExists {
			*paths += previousSumCount
		}
		pathSumsCache[nextPathSum]++

		helper(paths, node.Left, pathSumsCache, target, nextPathSum)
		helper(paths, node.Right, pathSumsCache, target, nextPathSum)

		pathSumsCache[nextPathSum]--
	}
}

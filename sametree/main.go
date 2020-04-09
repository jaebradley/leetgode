package sametree

/*
https://leetcode.com/problems/same-tree/

Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Example 1:

Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true
Example 2:

Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false
Example 3:

Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false

Approach:

* Like a lot of tree problems, the solution involves a recursive solution
* For a given node, it's the same as another node if they are both
	* nil
	* OR they are both non-nil and have the same value
		* AND their left subtrees are the same tree and their right subtrees are the same tree
* To check the left / right subtrees, can simply use the same function (hence the recursion)
*/

// TreeNode represents binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return (p == nil && q == nil) ||
		(p != nil && q != nil && p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right))
}

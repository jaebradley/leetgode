package validatebinarysearchtree

import "math"

/*
https://leetcode.com/problems/validate-binary-search-tree/

Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

Approach:

* Keep track of the range of possible values for a given node
* This range is determined by whether the parent was a left or a right branch
* If it was a left branch, all downstream branches must have values less than it, so range is previous min (at parent) and max is parent value
* If it was a right branch, all downstream branches must have values greater than it, so range is parent value (as the min) and max is previous max
*/

// TreeNode represents binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}

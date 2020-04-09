package diameterofbinarytree

/*
https://leetcode.com/problems/diameter-of-binary-tree/

Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

Example:
Given a binary tree

          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.

Approach:

* Recursive solution where the longest path at a given node is the longest path for the left node + the longest path for the right node + 1
* Leaf nodes have a longest path of 1
*/

// TreeNode is the definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter = 0
	calculateMaxHeight(root, &diameter)
	return diameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateMaxHeight(node *TreeNode, diameter *int) int {
	var maxHeight = 0

	if node != nil {
		var maxHeightLeftNode = calculateMaxHeight(node.Left, diameter)
		var maxHeightRightNode = calculateMaxHeight(node.Right, diameter)

		*diameter = max(*diameter, maxHeightLeftNode+maxHeightRightNode)
		maxHeight = max(maxHeightLeftNode, maxHeightRightNode) + 1
	}

	return maxHeight
}

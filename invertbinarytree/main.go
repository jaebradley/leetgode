package invertbinarytree

/*
https://leetcode.com/problems/invert-binary-tree/

Invert a binary tree.

Example:

Input:

     4
   /   \
  2     7
 / \   / \
1   3 6   9
Output:

     4
   /   \
  7     2
 / \   / \
9   6 3   1

Approach:

To invert a node, first invert it's children.
Then replace Left and Right.
Then return node.
Only invert children if node is non-nil.
*/

// TreeNode represents binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		invertTree(root.Left)
		invertTree(root.Right)

		temp := root.Left
		root.Left = root.Right
		root.Right = temp
	}

	return root
}

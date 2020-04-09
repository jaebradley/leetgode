package symmetrictree

/*
https://leetcode.com/problems/symmetric-tree/

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3


Follow up: Solve it both recursively and iteratively.

Recursive Approach:

* Need to compare the left and right subtrees from the root
* So basically, there's a pointer for the left and right subtree being evaluated at any given time (starting with the left child of the root and right child of root)
* Then the left child's left child should have the same subtree (just inverted) as the right child's right child
* Do the same with the left child's right child and right child's left child
* The "inversion" or symmetry of values is captured by comparing the current left node's left child to current right node's right child (same pattern for the remaining child)

*/

// TreeNode represents binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || areSymmetricTrees(root.Left, root.Right)
}

func areSymmetricTrees(firstTree *TreeNode, secondTree *TreeNode) bool {
	return (firstTree == nil && secondTree == nil) ||
		(firstTree != nil && secondTree != nil && firstTree.Val == secondTree.Val && areSymmetricTrees(firstTree.Left, secondTree.Right) && areSymmetricTrees(firstTree.Right, secondTree.Left))
}

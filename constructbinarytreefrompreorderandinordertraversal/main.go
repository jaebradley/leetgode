package constructbinarytreefrompreorderandinordertraversal

/*
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
	 15   7

Approach

* The first element in a pre-order is the root
* Everything left of that element in in-order is left and everything right of that element in in-order is right
* So in next (recursive) calculation, the left subtree's pre-order start (i.e. the left's root) is the previous pre-order index + 1
	* The inorder to consider is from previous inorder start to in-order root index - 1
* In next (recursive) calculation for right subtree, the pre-order start is the current pre-order start index + 1 + in-order root-index - in-order start-index
  * The inorder to consider is from the inorder root index + 1 to the inorder end index
*/

// TreeNode is the definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(0, 0, len(inorder)-1, preorder, inorder)
}

func helper(preorderStartIndex int, inorderStartIndex int, inorderedEndIndex int, preorder []int, inorder []int) *TreeNode {
	if preorderStartIndex > len(preorder)-1 || inorderStartIndex > inorderedEndIndex {
		return nil
	}

	rootValue := preorder[preorderStartIndex]
	inorderRootIndex := 0

	for index, value := range inorder {
		if value == rootValue {
			inorderRootIndex = index
			break
		}
	}

	return &TreeNode{
		Val:   rootValue,
		Left:  helper(preorderStartIndex+1, inorderStartIndex, inorderRootIndex-1, preorder, inorder),
		Right: helper(preorderStartIndex+inorderRootIndex-inorderStartIndex+1, inorderRootIndex+1, inorderedEndIndex, preorder, inorder),
	}
}

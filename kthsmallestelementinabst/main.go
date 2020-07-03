package kthsmallestelementinabst

/*
https://leetcode.com/problems/kth-smallest-element-in-a-bst/

Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.



Example 1:

Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1
Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
Output: 3
Follow up:
What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?



Constraints:

The number of elements of the BST is between 1 to 10^4.
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Approach:

* Push all left nodes to stack
* Pop off nodes
* For each popped off node, decrement count
* Once count is 0 return node value
* Get the right node (if it exists)
* Push right node
* Push all it's left nodes
*/

// TreeNode is the definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	nodes := []TreeNode{}

	for root != nil {
		nodes = append(nodes, *root)
		root = *&root.Left
	}

	for k > 0 {
		lastNode := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		k--

		if k == 0 {
			return lastNode.Val
		}

		nextNode := lastNode.Right
		for nextNode != nil {
			nodes = append(nodes, *nextNode)
			nextNode = nextNode.Left
		}
	}

	// shouldn't ever get here
	return -1
}

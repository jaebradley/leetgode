package verifypreorderserializationofabinarytree

import "strings"

/*
https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/

One way to serialize a binary tree is to use pre-order traversal. When we encounter a non-null node, we record the node's value. If it is a null node, we record using a sentinel value such as #.

     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
For example, the above binary tree can be serialized to the string "9,3,4,#,#,1,#,#,2,#,6,#,#", where # represents a null node.

Given a string of comma separated values, verify whether it is a correct preorder traversal serialization of a binary tree. Find an algorithm without reconstructing the tree.

Each comma separated value in the string must be either an integer or a character '#' representing null pointer.

You may assume that the input format is always valid, for example it could never contain two consecutive commas such as "1,,3".

Example 1:

Input: "9,3,4,#,#,1,#,#,2,#,6,#,#"
Output: true
Example 2:

Input: "1,#"
Output: false
Example 3:

Input: "9,#,#,1"
Output: false

Approach:

* Preorder traversal is root, left, right
* If null values are treated as leaves, then binary tree will be full
* A full binary tree is a tree in which every node other than the leaves has two children
* A full binary tree also has the property where # of leaves = # of non-leaves + 1
* So in order for serialization to be valid, the count of #'s must be equal to the count of non-# + 1
* Need to also account for a "valid" diff that is summed via an invalid combination like "#,#,2"
	* To handle this case, need to ensure diff is never 1 until the last node - in other words, never a case where the # of leaves is greater than number of non-leaves until last node

*/

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	diff := 0
	for _, node := range nodes {
		if diff == 1 {
			return false
		}
		if node == "#" {
			diff++
		} else {
			diff--
		}
	}

	return diff == 1
}

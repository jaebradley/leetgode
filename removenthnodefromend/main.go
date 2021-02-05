package removenthnodefromend

/**
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Follow up: Could you do this in one pass?

Example 1:


Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]

Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

Approach:

1. Two pointers - iterate through list using 1 pointer, n + 1 times
2. Start second pointer from beginning of list - when the first pointer gets to the end of the list, the second pointer will be at the node
BEFORE the nth node from the end of the list
3. Change the Next references for the current pointer (which is before the nth from end pointer) to skip a node
4. Using the dummy node reference return the Next that it refers to (this is to protect against when n is the size of the list).
In such a case, the dummy node would be the node before the nth node and thus would be the node that would "skip" the head node
*/

// ListNode is the definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var start = ListNode{0, head}
	var endNode = &start
	var priorToNthNodeFromEnd = &start

	for i := 0; i < n+1; i++ {
		endNode = endNode.Next
	}

	for endNode != nil {
		endNode = endNode.Next
		priorToNthNodeFromEnd = priorToNthNodeFromEnd.Next
	}

	priorToNthNodeFromEnd.Next = priorToNthNodeFromEnd.Next.Next

	return start.Next
}

package reverselinkedlist

/*
https://leetcode.com/problems/reverse-linked-list/

Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?

Recursive Approach:

Reversing a linked list at a given node is like reversing the list at the next node and just modifying the next pointer of the current node.
reverseList(1 -> 2 -> 3) => reverseList(3) -> 2 -> 1

Iterative Approach:

Keep track of previous node.
While current node is non-nil, swap current node with next node.
Current node's new next is previous node.
Make next node the current node.
Return previous node when current node becomes nil
*/

// ListNode is a node in a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func iterativelyReverseList(head *ListNode) *ListNode {
	var previous *ListNode = nil

	for head != nil {
		next := head.Next
		head.Next = previous
		previous = head
		head = next
	}

	return previous
}

func recursivelyReverseList(head *ListNode) *ListNode {
	return recursiveHelper(head, nil)
}

func recursiveHelper(currentHead *ListNode, newNext *ListNode) *ListNode {
	if currentHead == nil {
		return newNext
	}

	next := currentHead.Next
	currentHead.Next = newNext
	return recursiveHelper(next, currentHead)
}

package palindromelinkedlist

/*
https://leetcode.com/problems/palindrome-linked-list/

Given a singly linked list, determine if it is a palindrome.

Example 1:

Input: 1->2
Output: false
Example 2:

Input: 1->2->2->1
Output: true
Follow up:
Could you do it in O(n) time and O(1) space?
*/

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var previousNode *ListNode = nil
	currentNode := head

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}

	return previousNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	fast = head

	slow = reverseList(slow)

	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}

		slow = slow.Next
		fast = fast.Next
	}

	return true
}

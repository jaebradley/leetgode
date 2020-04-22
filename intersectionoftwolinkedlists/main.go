package intersectionoftwolinkedlists

/*
https://leetcode.com/problems/intersection-of-two-linked-lists/

Notes:

If the two linked lists have no intersection at all, return null.
The linked lists must retain their original structure after the function returns.
You may assume there are no cycles anywhere in the entire linked structure.
Your code should preferably run in O(n) time and use only O(1) memory.
*/

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pointerA := headA
	pointerB := headB

	for pointerA != pointerB {
		if pointerA == nil {
			pointerA = headB
		} else {
			pointerA = pointerA.Next
		}

		if pointerB == nil {
			pointerB = headA
		} else {
			pointerB = pointerB.Next
		}
	}

	return pointerA
}

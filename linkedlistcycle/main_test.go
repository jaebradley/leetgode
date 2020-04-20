package linkedlistcycle

import "testing"

func TestEmptyListDoesNotHaveCycle(t *testing.T) {
	if hasCycle(nil) != false {
		t.Fatalf("Expected empty list not to have cycle")
	}
}

func TestSingleElementListDoesNotHaveCycle(t *testing.T) {
	if hasCycle(&ListNode{Val: 1}) != false {
		t.Fatalf("Expected single element list not to have cycle")
	}
}

func TestDoubleElementListDoesNotHaveCycle(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	if hasCycle(list) != false {
		t.Fatalf("Expected double element list not to have cycle")
	}
}

func TestDoubleElementListCycle(t *testing.T) {
	var second ListNode

	head := &ListNode{
		Val:  1,
		Next: &second,
	}

	second = ListNode{
		Val:  2,
		Next: head,
	}

	if hasCycle(head) != true {
		t.Fatalf("Expected double element list that wraps around itself to have cycle")
	}
}

func TestTripleElementListCycle(t *testing.T) {
	var second ListNode
	var third ListNode

	head := &ListNode{
		Val:  1,
		Next: &second,
	}

	second = ListNode{
		Val:  2,
		Next: &third,
	}

	third = ListNode{
		Val:  3,
		Next: &second,
	}

	if hasCycle(head) != true {
		t.Fatalf("Expected triple element list that has cycle betwen second and third elements to return true")
	}
}

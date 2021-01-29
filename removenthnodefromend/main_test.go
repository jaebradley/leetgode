package removenthnodefromend

import "testing"

func TestSingleNodeList(t *testing.T) {
	updated := removeNthFromEnd(
		&ListNode{1, nil},
		1,
	)

	if updated != nil {
		t.Fatalf("Expected list to be gone")
	}
}

func TestWhenNIs1ForThreeElementList(t *testing.T) {
	updated := removeNthFromEnd(
		&ListNode{
			1,
			&ListNode{
				2,
				&ListNode{
					3,
					nil,
				},
			},
		},
		1,
	)

	if updated.Val != 1 {
		t.Fatalf("Expected head to be 1")
	}

	if updated.Next.Val != 2 {
		t.Fatalf("Expected second node to be 2")
	}

	if updated.Next.Next != nil {
		t.Fatalf("Expected second node to be last")
	}
}

func TestWhenNIsLengthOfThreeElementList(t *testing.T) {
	updated := removeNthFromEnd(
		&ListNode{
			1,
			&ListNode{
				2,
				&ListNode{
					3,
					nil,
				},
			},
		},
		3,
	)

	if updated.Val != 2 {
		t.Fatalf("Expected head to be 2")
	}

	if updated.Next.Val != 3 {
		t.Fatalf("Expected second node to be 3")
	}

	if updated.Next.Next != nil {
		t.Fatalf("Expected second node to be last")
	}
}

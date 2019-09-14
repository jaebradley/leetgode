package mergetwosortedlists

import "testing"

func TestTwoListsOfSameSize(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	l2 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}
	result := mergeTwoLists(l1, l2)
	if result.Val != 1 {
		t.Fatalf("Expected first list node to be 1")
	}

	if result.Next.Val != 2 {
		t.Fatalf("Expected second list node to be 2")
	}

	if result.Next.Next.Val != 3 {
		t.Fatalf("Expected third list node to be 3")
	}

	if result.Next.Next.Next.Val != 4 {
		t.Fatalf("Expected fourth list node to be 4")
	}
}

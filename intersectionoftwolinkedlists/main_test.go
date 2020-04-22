package intersectionoftwolinkedlists

import "testing"

func TestIntersectionOfListsWithSameLength(t *testing.T) {
	sharedHead := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 8,
			},
		},
	}

	first := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: sharedHead,
		},
	}

	second := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: sharedHead,
		},
	}

	if getIntersectionNode(first, second) != sharedHead {
		t.Fatalf("Intersection node is not expected shared node")
	}
}

func TestIntersectionOfListsWithDifferentLengths(t *testing.T) {
	sharedHead := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 8,
			},
		},
	}

	first := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: sharedHead,
			},
		},
	}

	second := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: sharedHead,
		},
	}

	intersection := getIntersectionNode(first, second)
	if intersection != sharedHead {
		t.Fatalf("Intersection node is not expected shared node %v", intersection)
	}
}


func TestNoIntersectionForTwoSeparateLists(t *testing.T) {
	first := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	second := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	if getIntersectionNode(first, second) != nil {
		t.Fatalf("Intersection should be nil for two separate lists")
	}
}
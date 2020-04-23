package palindromelinkedlist

import "testing"

func TestTwoNodePalindrome(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
		},
	}

	if isPalindrome(list) != true {
		t.Fatalf("Expected list to be a palindrome")
	}
}

func TestThreeNodePalindrome(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	if isPalindrome(list) != true {
		t.Fatalf("Expected list to be a palindrome")
	}
}

func TestFourNodePalindrome(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	if isPalindrome(list) != true {
		t.Fatalf("Expected list to be a palindrome")
	}
}

func TestTwoNodeInvalidPalindrome(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	if isPalindrome(list) != false {
		t.Fatalf("Expected list to be a palindrome")
	}
}

func TestThreeNodeInvalidPalindrome(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	if isPalindrome(list) != false {
		t.Fatalf("Expected list to be a palindrome")
	}
}

func TestFourNodeInvalidPalindrome(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	if isPalindrome(list) != false {
		t.Fatalf("Expected list to be a palindrome")
	}
}

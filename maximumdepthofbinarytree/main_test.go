package maximumdepthofbinarytree

import "testing"

func TestNilRoot(t *testing.T) {
	if maxDepth(nil) != 0 {
		t.Fatalf("Expected max depth of nil root to be 0")
	}
}

func TestNonNilRootWithoutChildren(t *testing.T) {
	root := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	if maxDepth(&root) != 1 {
		t.Fatalf("Expected max depth of 1 for non-nil root without children")
	}
}

func TestOnlySingleLeftChild(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}

	if maxDepth(&root) != 2 {
		t.Fatalf("Expected max depth of single left child to be 2")
	}
}

func TestOnlySingleRightChild(t *testing.T) {
	root := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	if maxDepth(&root) != 2 {
		t.Fatalf("Expected max depth of single left child to be 2")
	}
}

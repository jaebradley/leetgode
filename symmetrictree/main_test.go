package symmetrictree

import "testing"

func TestEmptyTree(t *testing.T) {
	if !isSymmetric(nil) {
		t.Fatalf("Expected empty tree to be symmetric")
	}
}

func TestSingleNode(t *testing.T) {
	if !isSymmetric(&TreeNode{Val: 1}) {
		t.Fatalf("Expected single node tree to be symmetric")
	}
}

func TestNodeWithLeftChild(t *testing.T) {
	if isSymmetric(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}) {
		t.Fatalf("Expected node with left child to not be symmetric")
	}
}

func TestNodeWithRightChild(t *testing.T) {
	if isSymmetric(&TreeNode{Val: 1, Right: &TreeNode{Val: 2}}) {
		t.Fatalf("Expected node with right child to not be symmetric")
	}
}

func TestNodeWithBothChildren(t *testing.T) {
	if !isSymmetric(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}) {
		t.Fatalf("Expected node with matching children to be symmetric")
	}
}

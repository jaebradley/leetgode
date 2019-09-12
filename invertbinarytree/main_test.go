package invertbinarytree

import "testing"

func TestInvertCompleteTwoLevelTree(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	result := invertTree(&root)
	if result.Val != 1 {
		t.Fatalf("Expected root to have a value of 1")
	}

	if result.Left.Val != 3 {
		t.Fatalf("Expected left value to have swapped to become 3")
	}

	if result.Right.Val != 2 {
		t.Fatalf("Expected right value to have swapped to become 2")
	}

	if result.Left.Left != nil {
		t.Fatalf("Expected left to have nil left child")
	}

	if result.Left.Right != nil {
		t.Fatalf("Expected left to have nil right child")
	}

	if result.Right.Left != nil {
		t.Fatalf("Expected right to have nil left child")
	}

	if result.Right.Right != nil {
		t.Fatalf("Expected right to have nil right child")
	}
}

func TestInvertTwoLevelTreeWithNilLeftChild(t *testing.T) {
	root := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	result := invertTree(&root)
	if result.Val != 1 {
		t.Fatalf("Expected root to have a value of 1")
	}

	if result.Left.Val != 3 {
		t.Fatalf("Expected left value to have swapped to become 3")
	}

	if result.Right != nil {
		t.Fatalf("Expected right value to have swapped to become nil")
	}

	if result.Left.Left != nil {
		t.Fatalf("Expected left to have nil left child")
	}

	if result.Left.Right != nil {
		t.Fatalf("Expected left to have nil right child")
	}
}

func TestInvertTwoLevelTreeWithoutARightChild(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	result := invertTree(&root)
	if result.Val != 1 {
		t.Fatalf("Expected root to have a value of 1")
	}

	if result.Right.Val != 2 {
		t.Fatalf("Expected left value to have swapped to become 3")
	}

	if result.Left != nil {
		t.Fatalf("Expected left value to have swapped to become nil")
	}

	if result.Right.Left != nil {
		t.Fatalf("Expected right to have nil left child")
	}

	if result.Right.Right != nil {
		t.Fatalf("Expected right to have nil right child")
	}
}

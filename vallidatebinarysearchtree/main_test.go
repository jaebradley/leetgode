package validatebinarysearchtree

import "testing"

func TestEmptyTreeReturnsTrue(t *testing.T) {
	if isValidBST(nil) != true {
		t.Fatalf("Expected empty root to be a valid BST")
	}
}

func TestTreeWithSingleRootNodeReturnsTrue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	if isValidBST(root) != true {
		t.Fatalf("Expected tree with single root node to be a valid BST")
	}
}

func TestTreeWitOnlyLeftChildWithLesserValueReturnsTrue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
	}
	if isValidBST(root) != true {
		t.Fatalf("Expected tree with single left child to be a valid BST")
	}
}

func TestTreeWitOnlyRightChildWithLargerValueReturnsTrue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	if isValidBST(root) != true {
		t.Fatalf("Expected tree with single right child to be a valid BST")
	}
}

func TestTreeWitOnlyLeftChildWithGreaterValueReturnsFalse(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	if isValidBST(root) != false {
		t.Fatalf("Expected tree with single left child with greater value to be an invalid BST")
	}
}

func TestTreeWitOnlyRightChildWithLesserValueReturnsTrue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 0,
		},
	}
	if isValidBST(root) != false {
		t.Fatalf("Expected tree with single right child with lesser value to be an invalid BST")
	}
}

func TestTreeWitOnlyLeftChildWithEqualValueReturnsFalse(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}
	if isValidBST(root) != false {
		t.Fatalf("Expected tree with single left child with equal value to be an invalid BST")
	}
}

func TestTreeWitOnlyRightChildWithEqualValueReturnsTrue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
		},
	}
	if isValidBST(root) != false {
		t.Fatalf("Expected tree with single right child with equal value to be an invalid BST")
	}
}

func TestTreeWithNestedLeftNodeWhereRightNodeIsGreaterThanParentReturnsFalse(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 8,
			Right: &TreeNode{
				Val: 11,
			},
		},
	}
	if isValidBST(root) != false {
		t.Fatalf("Expected tree with nested left subtree with a right tree greater than parent is an invalid BST")
	}
}

func TestTreeWithNestedLeftNodeWhereRightNodeIsLessThanParentReturnsTrue(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 8,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	if isValidBST(root) != true {
		t.Fatalf("Expected tree with nested left subtree with a right tree less than parent is a valid BST")
	}
}

func TestTreeWithNestedRightNodeWhereLeftNodeIsLessThanParentReturnsFalse(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Right: &TreeNode{
			Val: 12,
			Left: &TreeNode{
				Val: 9,
			},
		},
	}
	if isValidBST(root) != false {
		t.Fatalf("Expected tree with nested right subtree with a left tree less than parent is an invalid BST")
	}
}

func TestTreeWithNestedRightNodeWhereRightNodeIsNotGreaterThanParentButGreaterThanGrandparentReturnsFalse(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 11,
			},
		},
	}
	if isValidBST(root) != false {
		t.Fatalf("Expected tree with nested right subtree with a right tree less than parent is an invalid BST")
	}
}

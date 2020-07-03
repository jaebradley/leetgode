package kthsmallestelementinabst

import "testing"

func TestOnlyRootReturnsRootValue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	if kthSmallest(root, 1) != 1 {
		t.Errorf("Expected kth smallest value to be 1")
	}
}

func TestRootAndRightChildReturnsRightChildValueWhenGetting2ndSmallestValue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	if kthSmallest(root, 2) != 2 {
		t.Errorf("Expected kth smallest value to be 2")
	}
}

func TestRootAndRightChildReturnsLeftChildValueWhenGettingSmallestValue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	if kthSmallest(root, 1) != 1 {
		t.Errorf("Expected kth smallest value to be 1")
	}
}

func TestRootAndLeftChildReturnsLeftChildValueWhenGettingSmallestValue(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}
	if kthSmallest(root, 1) != 1 {
		t.Errorf("Expected kth smallest value to be 1")
	}
}

func TestRootAndLeftChildReturnsLeftChildValueWhenGetting2ndSmallestValue(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}
	if kthSmallest(root, 2) != 2 {
		t.Errorf("Expected kth smallest value to be 2")
	}
}

func TestRootAndLeftAndRightChildReturnsLeftChildValueWhenGettingSmallestValue(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	if kthSmallest(root, 1) != 1 {
		t.Errorf("Expected kth smallest value to be 1")
	}
}

func TestRootAndLeftAndRightChildReturnsLeftChildValueWhenGetting2ndSmallestValue(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	if kthSmallest(root, 2) != 2 {
		t.Errorf("Expected kth smallest value to be 2")
	}
}

func TestRootAndLeftAndRightChildReturnsLeftChildValueWhenGetting3rdSmallestValue(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	if kthSmallest(root, 3) != 3 {
		t.Errorf("Expected kth smallest value to be 3")
	}
}

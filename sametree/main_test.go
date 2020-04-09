package sametree

import "testing"

func TestTwoEmptyTrees(t *testing.T) {
	if !isSameTree(nil, nil) {
		t.Fatalf("Expected two empty trees to be the same")
	}
}

func TestEmptyTreeAndExistingTree(t *testing.T) {
	if isSameTree(nil, &TreeNode{Val: 1}) {
		t.Fatalf("Expected empty tree not to be same as non-empty tree")
	}
}

func TestTreesWithExistingLeftDescendant(t *testing.T) {
	var firstTree = TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	var secondTree = TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	if !isSameTree(&firstTree, &secondTree) {
		t.Fatalf("Expected trees with the same existing left desendants to be equal")
	}
}

func TestTreesWithExistingRightDescendant(t *testing.T) {
	var firstTree = TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	var secondTree = TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	if !isSameTree(&firstTree, &secondTree) {
		t.Fatalf("Expected trees with the same existing right desendants to be equal")
	}
}

func TestTreesWithExistingDescendants(t *testing.T) {
	var firstTree = TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	var secondTree = TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	if !isSameTree(&firstTree, &secondTree) {
		t.Fatalf("Expected trees with the same existing desendants to be equal")
	}
}

package lowestcommonancestorofbinarysearchtree

import "testing"

func TestLeftChild(t *testing.T) {
	q := TreeNode{0, nil, nil}
	tree := TreeNode{1, &q, nil}
	p := tree

	if &tree != lowestCommonAncestor(&tree, &p, &q) {
		t.Fatalf("Expected root to be LCA of left node and root")
	}
}

func TestRightChild(t *testing.T) {
	q := TreeNode{2, nil, nil}
	tree := TreeNode{1, nil, &q}
	p := tree

	if &tree != lowestCommonAncestor(&tree, &p, &q) {
		t.Fatalf("Expected root to be LCA of right node and root")
	}
}

func TestLeftAndRightChildShareRootAsLCA(t *testing.T) {
	p := TreeNode{0, nil, nil}
	q := TreeNode{2, nil, nil}
	tree := TreeNode{1, &p, &q}

	if &tree != lowestCommonAncestor(&tree, &p, &q) {
		t.Fatalf("Expected root to be LCA of left node and right node")
	}
}

func TestLeftAndRightChildrenAtDifferentLevelsShareRootAsLCA(t *testing.T) {
	p := TreeNode{0, nil, nil}
	q := TreeNode{2, nil, nil}
	tree := TreeNode{
		1,
		&p,
		&TreeNode{3, &q, nil},
	}

	if &tree != lowestCommonAncestor(&tree, &p, &q) {
		t.Fatalf("Expected root to be LCA of left node and right node")
	}
}

package diameterofbinarytree

import "testing"

func TestEmptyTree(t *testing.T) {
	if diameterOfBinaryTree(nil) != 0 {
		t.Fatalf("Expected diameter of empty tree to be 0")
	}
}

func TestTreeWithNoChildren(t *testing.T) {
	if (diameterOfBinaryTree(&TreeNode{Val: 1})) != 0 {
		t.Fatalf("Expected diameter of single node tree to be 0")
	}
}

func TestTreeWithOnlyLeftChild(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	if diameterOfBinaryTree(&root) != 1 {
		t.Fatalf("Expected diameter of tree with single left child to be 1")
	}
}

func TestTreeWithOnlyRightChild(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	if diameterOfBinaryTree(&root) != 1 {
		t.Fatalf("Expected diameter of tree with single right child to be 1")
	}
}

func TestTreeWithLeftAndRightChild(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	if diameterOfBinaryTree(&root) != 2 {
		t.Fatalf("Expected diameter of tree with both left and right child to be 2")
	}
}

func TestTreeWithTwoLevelsOfChildrenOnLeftAndOneLevelOfChildrenOnRight(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	if diameterOfBinaryTree(&root) != 3 {
		t.Fatalf("Expected diameter of tree with two levels of children on left and one level on right to be 3")
	}
}

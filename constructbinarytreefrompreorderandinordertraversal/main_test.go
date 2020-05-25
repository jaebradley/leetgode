package constructbinarytreefrompreorderandinordertraversal

import "testing"

func TestSingleLeftChildWithRightSubtree(t *testing.T) {
	actual := buildTree(
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
	)

	if actual.Val != 3 {
		t.Fatalf("Expected root to have value of 3")
	}

	if actual.Left.Val != 9 {
		t.Fatalf("Expected left to have value of 9")
	}

	if actual.Right.Val != 20 {
		t.Fatalf("Expected right to have value of 20")
	}

	if actual.Right.Left.Val != 15 {
		t.Fatalf("Expected left child of right child to have value of 15")
	}

	if actual.Right.Right.Val != 7 {
		t.Fatalf("Expected right child of right child to have value of 7 ")
	}
}

func TestEmptyTree(t *testing.T) {
	actual := buildTree([]int{}, []int{})

	if actual != nil {
		t.Fatalf("Expected results to be empty")
	}
}

func TestOnlyRoot(t *testing.T) {
	actual := buildTree([]int{1}, []int{1})

	if actual.Val != 1 {
		t.Fatalf("Expected root value to be 1")
	}

	if actual.Right != nil {
		t.Fatalf("Expected root to not have right child")
	}

	if actual.Left != nil {
		t.Fatalf("Expected root to not have left child")
	}
}

func TestOnlyRootAndLeftChild(t *testing.T) {
	actual := buildTree([]int{1, 2}, []int{2, 1})

	if actual.Val != 1 {
		t.Fatalf("Expected root value to be 1")
	}

	if actual.Right != nil {
		t.Fatalf("Expected root to not have right child")
	}

	if actual.Left.Val != 2 {
		t.Fatalf("Expected left value to be 2")
	}
}

func TestOnlyRootAndRightChild(t *testing.T) {
	actual := buildTree([]int{1, 2}, []int{1, 2})

	if actual.Val != 1 {
		t.Fatalf("Expected root value to be 1")
	}

	if actual.Left != nil {
		t.Fatalf("Expected root to not have left child")
	}

	if actual.Right.Val != 2 {
		t.Fatalf("Expected right value to be 2")
	}
}

func TestLeftChildrenSingleRight(t *testing.T) {
	actual := buildTree([]int{3, 1, 2, 4}, []int{1, 2, 3, 4})

	if actual.Val != 3 {
		t.Fatalf("Expected root value to be 3")
	}

	if actual.Left.Val != 1 {
		t.Fatalf("Expected left child value to be 1")
	}

	if actual.Left.Right.Val != 2 {
		t.Fatalf("Expected left child's right child to be 2")
	}

	if actual.Left.Left != nil {
		t.Fatalf("Expected left child's right child to be empty")
	}

	if actual.Right.Val != 4 {
		t.Fatalf("Expected rigth child's value to be 4")
	}
}

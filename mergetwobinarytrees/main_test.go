package mergetwobinarytrees

import "testing"

func TestSummingTwoLevelCompleteTree(t *testing.T) {
	t1 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	t2 := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	actualResult := mergeTrees(&t1, &t2)
	if actualResult.Val != 3 {
		t.Fatalf("Expected root to have sum of 3")
	}

	if actualResult.Left.Val != 4 {
		t.Fatalf("Expected left to have sum of 4")
	}

	if actualResult.Right.Val != 5 {
		t.Fatalf("Expected right to have sum of 5")
	}

	if actualResult.Left.Left != nil {
		t.Fatalf("Expected left to have nil left child")
	}

	if actualResult.Left.Right != nil {
		t.Fatalf("Expected left to have nil right child")
	}

	if actualResult.Right.Left != nil {
		t.Fatalf("Expected right to have nil left child")
	}

	if actualResult.Right.Right != nil {
		t.Fatalf("Expected right to have nil right child")
	}
}

func TestSummingTwoLevelLeftEmptyTree(t *testing.T) {
	t1 := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	t2 := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	actualResult := mergeTrees(&t1, &t2)
	if actualResult.Val != 3 {
		t.Fatalf("Expected root to have sum of 3")
	}

	if actualResult.Left.Val != 1 {
		t.Fatalf("Expected left to have sum of 4")
	}

	if actualResult.Right.Val != 5 {
		t.Fatalf("Expected right to have sum of 5")
	}

	if actualResult.Left.Left != nil {
		t.Fatalf("Expected left to have nil left child")
	}

	if actualResult.Left.Right != nil {
		t.Fatalf("Expected left to have nil right child")
	}

	if actualResult.Right.Left != nil {
		t.Fatalf("Expected right to have nil left child")
	}

	if actualResult.Right.Right != nil {
		t.Fatalf("Expected right to have nil right child")
	}
}

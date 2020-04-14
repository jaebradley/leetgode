package pathsumthree

import "testing"

func TestEmptyTreeHasZeroPaths(t *testing.T) {
	if pathSum(nil, 0) != 0 {
		t.Fatalf("Expected empty tree sum to be 0")
	}
}

func TestSingleTreeHasSinglePath(t *testing.T) {
	if pathSum(&TreeNode{Val: 1}, 1) != 1 {
		t.Fatalf("Expected path sum to be 1")
	}
}

func TestLeftMatchingSumFromRoot(t *testing.T) {
	sum := pathSum(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
		3,
	)
	if sum != 1 {
		t.Fatalf("Expected path sum to be 1 instead of %d", sum)
	}
}

func TestRightMatchingSumFromRoot(t *testing.T) {
	sum := pathSum(
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		3,
	)
	if sum != 1 {
		t.Fatalf("Expected path sum to be 1 instead of %d", sum)
	}
}

func TestDirectChildrenMatchingSumFromRoot(t *testing.T) {
	sum := pathSum(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		3,
	)
	if sum != 2 {
		t.Fatalf("Expected path sum to be 2 instead of %d", sum)
	}
}

func TestLeftSubchildMatchesTarget(t *testing.T) {
	sum := pathSum(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
				},
			},
		},
		8,
	)
	if sum != 1 {
		t.Fatalf("Expected path sum to be 1 instead of %d", sum)
	}
}

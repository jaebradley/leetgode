package minimumpathsum

import "testing"

func TestSingleCell(t *testing.T) {
	grid := [][]int{{1}}
	calculated := minPathSum(grid)

	if calculated != 1 {
		t.Fatalf("Expected min path sum to be 1 instead of %d", calculated)
	}
}

func TestSingleRow(t *testing.T) {
	grid := [][]int{
		{1, 2, 3},
	}

	calculated := minPathSum(grid)

	if calculated != 6 {
		t.Fatalf("Expected min path sum to be 6 instead of %d", calculated)
	}
}

func TestSingleColumn(t *testing.T) {
	grid := [][]int{
		{1},
		{2},
		{3},
	}

	calculated := minPathSum(grid)

	if calculated != 6 {
		t.Fatalf("Expected min path sum to be 6 instead of %d", calculated)
	}
}

func TestTwoByTwo(t *testing.T) {
	grid := [][]int{
		{1, 2},
		{3, 4},
	}

	calculated := minPathSum(grid)

	if calculated != 7 {
		t.Fatalf("Expected min path sum to be 7 instead of %d", calculated)
	}
}

func TestTwoByThree(t *testing.T) {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	calculated := minPathSum(grid)

	if calculated != 12 {
		t.Fatalf("Expected min path sum to be 12 instead of %d", calculated)
	}
}

func TestThreeByTwo(t *testing.T) {
	grid := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	calculated := minPathSum(grid)

	if calculated != 13 {
		t.Fatalf("Expected min path sum to be 13 instead of %d", calculated)
	}
}

func TestThreeByThree(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	calculated := minPathSum(grid)

	if calculated != 7 {
		t.Fatalf("Expected min path sum to be 7 instead of %d", calculated)
	}
}

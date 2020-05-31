package uniquepaths

import "testing"

func TestSingleCell(t *testing.T) {
	paths := uniquePaths(1, 1)

	if paths != 1 {
		t.Fatalf("Expected paths to be 1 instead of %d", paths)
	}
}

func TestRow(t *testing.T) {
	paths := uniquePaths(10, 1)

	if paths != 1 {
		t.Fatalf("Expected paths to be 1 instead of %d", paths)
	}
}

func TestColumn(t *testing.T) {
	paths := uniquePaths(1, 10)

	if paths != 1 {
		t.Fatalf("Expected paths to be 1 instead of %d", paths)
	}
}

func TestTwoByTwo(t *testing.T) {
	paths := uniquePaths(2, 2)

	if paths != 2 {
		t.Fatalf("Expected paths to be 2 instead of %d", paths)
	}
}

func TestThreeByThree(t *testing.T) {
	paths := uniquePaths(3, 3)

	if paths != 6 {
		t.Fatalf("Expected paths to be 6 instead of %d", paths)
	}
}

func TestTwoByThree(t *testing.T) {
	paths := uniquePaths(2, 3)

	if paths != 3 {
		t.Fatalf("Expected paths to be 3 instead of %d", paths)
	}
}

func TestThreeByTwo(t *testing.T) {
	paths := uniquePaths(3, 2)

	if paths != 3 {
		t.Fatalf("Expected paths to be 3 instead of %d", paths)
	}
}

package findallnumbersdisappearedinanarray

import (
	"testing"
)

func TestArray(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := findDisappearedNumbers(nums)

	if result[0] != 5 {
		t.Fatalf("Expected 5 to not be in array")
	}

	if result[1] != 6 {
		t.Fatalf("Expected 6 to not be in array")
	}
}

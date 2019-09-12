package singlenumber

import "testing"

func TestThreeValueArray(t *testing.T) {
	nums := []int{1, 2, 1}
	if singleNumber(nums) != 2 {
		t.Fatalf("Expected [1, 2, 1] to return 2")
	}
}

package movezeroes

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var nums = []int{0, 1, 0, 3, 12}

	moveZeroes(nums)

	var expected = []int{1, 3, 12, 0, 0}
	for i, v := range nums {
		if v != expected[i] {
			t.Fatalf(fmt.Sprintf("Expected array to be modified in-place to %v and not %v", expected, nums))
		}
	}
}

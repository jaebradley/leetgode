package besttimetobuyandsellstock

import (
	"testing"
)

func TestArray(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(nums)

	if result != 5 {
		t.Fatalf("Expected max profit to be 5")
	}
}

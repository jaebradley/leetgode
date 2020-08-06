package coinchange

import "testing"

func TestZeroAmountReturnsZero(t *testing.T) {
	change := coinChange([]int{1}, 0)

	if change != 0 {
		t.Errorf("Expected change to be 0 instead of %d", change)
	}
}

func TestSingleCoinValueThatDoesNotMatchAmountReturnsNegativeOne(t *testing.T) {
	change := coinChange([]int{2}, 3)

	if change != -1 {
		t.Errorf("Expected change to be -1 instead of %d", change)
	}
}

func TestChooseTwoOfTheSameCoin(t *testing.T) {
	change := coinChange([]int{1, 5}, 11)

	if change != 3 {
		t.Errorf("Expected change to be 3 instead of %d", change)
	}
}

func TestChooseCoinValue(t *testing.T) {
	change := coinChange([]int{1, 2}, 2)

	if change != 1 {
		t.Errorf("Expected change to be 1 instead of %d", change)
	}
}

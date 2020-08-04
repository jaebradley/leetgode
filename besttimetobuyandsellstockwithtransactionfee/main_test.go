package besttimetobuyandsellstockwithtransactionfee

import "testing"

func TestZeroMaxProfit(t *testing.T) {
	profit := maxProfit([]int{1, 1, 1}, 10)
	if profit != 0 {
		t.Errorf("Expected profit to be 0 instead of %d", profit)
	}
}

func TestBuyingTwiceInSession(t *testing.T) {
	profit := maxProfit([]int{1, 2, 3, 8, 4, 9}, 2)
	if profit != 8 {
		t.Errorf("Expected profit to be 8 instead of %d", profit)
	}
}

package minimumcostfortickets

import "testing"

func TestDoublePurchase(t *testing.T) {
	minCost := mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15})
	if minCost != 11 {
		t.Fatalf("Expected minimum cost to be 11 and not %d", minCost)
	}
}

func TestUse7DayInsteadOf1Day(t *testing.T) {
	minCost := mincostTickets([]int{1, 2, 3, 4}, []int{2, 7, 15})
	if minCost != 7 {
		t.Fatalf("Expected minimum cost to be 7 and not %d", minCost)
	}
}

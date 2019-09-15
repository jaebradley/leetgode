package houserobber

import "testing"

func TestEveryOther(t *testing.T) {
	if rob([]int{1, 2, 3, 1}) != 4 {
		t.Fatalf("Expected max amount robbed to be 4")
	}
}

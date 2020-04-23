package shortestunsortedcontinuoussubarray

import "testing"

func TestZeroForAlreadySortedArray(t *testing.T) {
	length := findUnsortedSubarray([]int{1, 2, 3, 4, 5, 6})
	if length != 0 {
		t.Fatalf("Expected unsorted subarray to have length 0 for already sorted array instead of : %d", length)
	}
}

func TestFiveElementUnsortedSubarray(t *testing.T) {
	length := findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15})
	if length != 5 {
		t.Fatalf("Expected unsorted subarray to have length 5 instead of : %d", length)
	}
}

func TestWhenElementIncreasesAtEndButIsLessThanMaximum(t *testing.T) {
	length := findUnsortedSubarray([]int{2, 10, 9, 8, 7, 6, 7})
	if length != 6 {
		t.Fatalf("Expected unsorted subarray to have length 6 instead of : %d", length)
	}
}

func TestEmptyArrayIsZero(t *testing.T) {
	length := findUnsortedSubarray([]int{})
	if length != 0 {
		t.Fatalf("Expected unsorted subarray to have length 0 instead of : %d", length)
	}
}

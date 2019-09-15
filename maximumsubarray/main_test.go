package maximumsubarray

import "testing"

func TestNilSubarray(t *testing.T) {
	if maxSubArray(nil) != 0 {
		t.Fatalf("Expected max subarray of nil to be 0")
	}
}

func TestSingleElementSubarray(t *testing.T) {
	values := []int{10}
	if maxSubArray(values) != 10 {
		t.Fatalf("Expected max subarray of single array with 10 to be 10")
	}
}

func TestNegativeSingleElementInput(t *testing.T) {
	values := []int{-10}
	if maxSubArray(values) != -10 {
		t.Fatalf("Expected max subarray of single array with -10 to be -10")
	}
}

func TestMaxSubarrayWithNegatives(t *testing.T) {
	values := []int{4, -1, 2, 1}
	if maxSubArray(values) != 6 {
		t.Fatalf("Expected max subarray to be 6")
	}
}

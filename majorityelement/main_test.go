package majorityelement

import "testing"

func TestSingleElement(t *testing.T) {
	if majorityElement([]int{3}) != 3 {
		t.Fatalf("Expected majority element to be 3")
	}
}

func TestTwoElements(t *testing.T) {
	if majorityElement([]int{3, 3}) != 3 {
		t.Fatalf("Expected majority element to be 3")
	}
}

func TestThreeElements(t *testing.T) {
	if majorityElement([]int{3, 3, 4}) != 3 {
		t.Fatalf("Expected majority element to be 3")
	}
}

func TestFourElements(t *testing.T) {
	if majorityElement([]int{3, 2, 3, 3}) != 3 {
		t.Fatalf("Expected majority element to be 3")
	}
}

func TestFiveElements(t *testing.T) {
	if majorityElement([]int{2, 2, 3, 3, 3}) != 3 {
		t.Fatalf("Expected majority element to be 3")
	}
}

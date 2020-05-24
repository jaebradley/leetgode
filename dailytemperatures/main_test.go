package dailytemperatures

import (
	"testing"
)

func TestSortedAscending(t *testing.T) {
	temps := []int{1, 2, 3, 4}
	expected := []int{1, 1, 1, 0}
	output := dailyTemperatures(temps)
	if len(output) != 4 {
		t.Fatalf("Expected output to have length of 4 and not %d", len(output))
	}
	for index, wait := range output {
		expectedValue := expected[index]

		if wait != expectedValue {
			t.Fatalf("Expected wait: %d to be: %d at index: %d", wait, expectedValue, index)
		}
	}
}

func TestSortedDescending(t *testing.T) {
	temps := []int{4, 3, 2, 1}
	expected := []int{0, 0, 0, 0}
	output := dailyTemperatures(temps)
	if len(output) != 4 {
		t.Fatalf("Expected output to have length of 4 and not %d", len(output))
	}
	for index, wait := range output {
		expectedValue := expected[index]

		if wait != expectedValue {
			t.Fatalf("Expected wait: %d to be: %d at index: %d", wait, expectedValue, index)
		}
	}
}

func TestCombination(t *testing.T) {
	temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	expected := []int{1, 1, 4, 2, 1, 1, 0, 0}
	output := dailyTemperatures(temps)
	if len(output) != 8 {
		t.Fatalf("Expected output to have length of 8 and not %d", len(output))
	}
	for index, wait := range output {
		expectedValue := expected[index]

		if wait != expectedValue {
			t.Fatalf("Expected wait: %d to be: %d at index: %d", wait, expectedValue, index)
		}
	}
}

func TestEmpty(t *testing.T) {
	temps := []int{}
	expected := []int{}
	output := dailyTemperatures(temps)
	if len(output) != 0 {
		t.Fatalf("Expected output to have length of 0 and not %d", len(output))
	}
	for index, wait := range output {
		expectedValue := expected[index]

		if wait != expectedValue {
			t.Fatalf("Expected wait: %d to be: %d at index: %d", wait, expectedValue, index)
		}
	}
}

func TestSingle(t *testing.T) {
	temps := []int{1}
	expected := []int{0}
	output := dailyTemperatures(temps)

	if len(output) != 1 {
		t.Fatalf("Expected output to have length of 1 and not %d", len(output))
	}

	for index, wait := range output {
		expectedValue := expected[index]

		if wait != expectedValue {
			t.Fatalf("Expected wait: %d to be: %d at index: %d", wait, expectedValue, index)
		}
	}
}

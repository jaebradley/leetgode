package insertinterval

import "testing"

func TestInsertingWithoutOverlap(t *testing.T) {
	intervals := [][]int{
		{1, 2},
		{3, 4},
		{7, 8},
	}

	newInterval := []int{5, 6}
	expected := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	result := insert(intervals, newInterval)
	for i, value := range result {
		if value[0] != expected[i][0] || value[1] != expected[i][1] {
			t.Errorf("Expected %v instead of %v at index %d", expected[i], value, i)
		}
	}
}

func TestInsertingWithSingleOverlapFromStartOfExistingInterval(t *testing.T) {
	intervals := [][]int{
		{1, 2},
		{3, 4},
		{7, 8},
	}

	newInterval := []int{3, 5}
	expected := [][]int{
		{1, 2},
		{3, 5},
		{7, 8},
	}
	result := insert(intervals, newInterval)
	for i, value := range result {
		if value[0] != expected[i][0] || value[1] != expected[i][1] {
			t.Errorf("Expected %v instead of %v at index %d", expected[i], value, i)
		}
	}
}

func TestInsertingWithSingleOverlapFromEndOfExistingInterval(t *testing.T) {
	intervals := [][]int{
		{1, 2},
		{3, 4},
		{7, 8},
	}

	newInterval := []int{4, 5}
	expected := [][]int{
		{1, 2},
		{3, 5},
		{7, 8},
	}
	result := insert(intervals, newInterval)
	for i, value := range result {
		if value[0] != expected[i][0] || value[1] != expected[i][1] {
			t.Errorf("Expected %v instead of %v at index %d", expected[i], value, i)
		}
	}
}

func TestInsertingWithMultipleOverlapFromStartOfExistingInterval(t *testing.T) {
	intervals := [][]int{
		{1, 2},
		{3, 4},
		{7, 8},
	}

	newInterval := []int{3, 7}
	expected := [][]int{
		{1, 2},
		{3, 8},
	}
	result := insert(intervals, newInterval)
	for i, value := range result {
		if value[0] != expected[i][0] || value[1] != expected[i][1] {
			t.Errorf("Expected %v instead of %v at index %d", expected[i], value, i)
		}
	}
}

func TestInsertingWithMultipleOverlapFromEndOfExistingInterval(t *testing.T) {
	intervals := [][]int{
		{1, 2},
		{3, 4},
		{7, 8},
	}

	newInterval := []int{4, 7}
	expected := [][]int{
		{1, 2},
		{3, 8},
	}
	result := insert(intervals, newInterval)
	for i, value := range result {
		if value[0] != expected[i][0] || value[1] != expected[i][1] {
			t.Errorf("Expected %v instead of %v at index %d", expected[i], value, i)
		}
	}
}

func TestInsertingWithMultipleOverlapExtendingEndIndex(t *testing.T) {
	intervals := [][]int{
		{1, 2},
		{3, 4},
		{7, 8},
	}

	newInterval := []int{3, 9}
	expected := [][]int{
		{1, 2},
		{3, 9},
	}
	result := insert(intervals, newInterval)
	for i, value := range result {
		if value[0] != expected[i][0] || value[1] != expected[i][1] {
			t.Errorf("Expected %v instead of %v at index %d", expected[i], value, i)
		}
	}
}
func TestInsertingWithMultipleOverlapWithFollowingIntervalsExtendingEndIndex(t *testing.T) {
	intervals := [][]int{
		{1, 2},
		{3, 4},
		{7, 8},
		{100, 111},
	}

	newInterval := []int{3, 9}
	expected := [][]int{
		{1, 2},
		{3, 9},
		{100, 111},
	}
	result := insert(intervals, newInterval)
	for i, value := range result {
		if value[0] != expected[i][0] || value[1] != expected[i][1] {
			t.Errorf("Expected %v instead of %v at index %d", expected[i], value, i)
		}
	}
}

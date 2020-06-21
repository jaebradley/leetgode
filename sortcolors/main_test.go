package sortcolors

import "testing"

func TestSorted(t *testing.T) {
	nums := []int{0, 1, 2}
	expected := []int{0, 1, 2}

	sortColors(nums)

	for i := 0; i < len(nums); i++ {
		if expected[i] != nums[i] {
			t.Errorf("Expected value at index %d to be %d instead of %d", i, expected[i], nums[i])
		}
	}
}

func TestReversed(t *testing.T) {
	nums := []int{2, 1, 0}
	expected := []int{0, 1, 2}

	sortColors(nums)

	for i := 0; i < len(nums); i++ {
		if expected[i] != nums[i] {
			t.Errorf("Expected value at index %d to be %d instead of %d", i, expected[i], nums[i])
		}
	}
}

func OnlyZeros(t *testing.T) {
	nums := []int{0, 0, 0}
	expected := []int{0, 0, 0}

	sortColors(nums)

	for i := 0; i < len(nums); i++ {
		if expected[i] != nums[i] {
			t.Errorf("Expected value at index %d to be %d instead of %d", i, expected[i], nums[i])
		}
	}
}

func OnlyOnes(t *testing.T) {
	nums := []int{1, 1, 1}
	expected := []int{1, 1, 1}

	sortColors(nums)

	for i := 0; i < len(nums); i++ {
		if expected[i] != nums[i] {
			t.Errorf("Expected value at index %d to be %d instead of %d", i, expected[i], nums[i])
		}
	}
}

func OnlyTwos(t *testing.T) {
	nums := []int{2, 2, 2}
	expected := []int{2, 2, 2}

	sortColors(nums)

	for i := 0; i < len(nums); i++ {
		if expected[i] != nums[i] {
			t.Errorf("Expected value at index %d to be %d instead of %d", i, expected[i], nums[i])
		}
	}
}

func OneAfterTwo(t *testing.T) {
	nums := []int{2, 0, 1}
	expected := []int{0, 1, 2}

	sortColors(nums)

	for i := 0; i < len(nums); i++ {
		if expected[i] != nums[i] {
			t.Errorf("Expected value at index %d to be %d instead of %d", i, expected[i], nums[i])
		}
	}
}

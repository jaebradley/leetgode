package houserobber2

import "testing"

func TestEmptyValuesReturnsZero(t *testing.T) {
	maxValue := rob([]int{})

	if maxValue != 0 {
		t.Errorf("Expected maximum value to be 0 and not %d", maxValue)
	}
}

func TestSingleValueReturnsValue(t *testing.T) {
	maxValue := rob([]int{1})

	if maxValue != 1 {
		t.Errorf("Expected maximum value to be 1 and not %d", maxValue)
	}
}

func TestThreeValueReturnsMiddleValueWhenItIsGreatest(t *testing.T) {
	maxValue := rob([]int{2, 3, 2})

	if maxValue != 3 {
		t.Errorf("Expected maximum value to be 3 and not %d", maxValue)
	}
}

func TestThreeValueReturnsFirstValueWhenItIsGreatest(t *testing.T) {
	maxValue := rob([]int{3, 2, 2})

	if maxValue != 3 {
		t.Errorf("Expected maximum value to be 3 and not %d", maxValue)
	}
}

func TestThreeValueReturnsLastValueWhenItIsGreatest(t *testing.T) {
	maxValue := rob([]int{3, 2, 3})

	if maxValue != 3 {
		t.Errorf("Expected maximum value to be 3 and not %d", maxValue)
	}
}

func TestFourValuesFirstAndThirdValueSum(t *testing.T) {
	maxValue := rob([]int{1, 2, 3, 1})

	if maxValue != 4 {
		t.Errorf("Expected maximum value to be 4 and not %d", maxValue)
	}
}

func TestFourValuesSecondAndFourthValueSum(t *testing.T) {
	maxValue := rob([]int{1, 2, 3, 3})

	if maxValue != 5 {
		t.Errorf("Expected maximum value to be 5 and not %d", maxValue)
	}
}

func TestFiveValuesSkipRobbingOneTimeSum(t *testing.T) {
	maxValue := rob([]int{10, 1, 1, 10, 1})

	if maxValue != 20 {
		t.Errorf("Expected maximum value to be 20 and not %d", maxValue)
	}
}

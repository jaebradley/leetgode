package courseschedule

import "testing"

func TestIndependentSinglePrerequisite(t *testing.T) {
	prerequisites := [][]int{
		{1, 0},
	}

	if canFinish(1, prerequisites) != true {
		t.Error("Expected independent prerequisites to be able to finish")
	}
}

func TestDependentPrerequisites(t *testing.T) {
	prerequisites := [][]int{
		{1, 0},
		{0, 1},
	}

	if canFinish(2, prerequisites) != false {
		t.Error("Expected dependent prerequisites not to be able to finish")
	}
}

func TestDependentPrerequisitesWhenNumberOfCoursesGreaterThanPrerequisitesLength(t *testing.T) {
	prerequisites := [][]int{
		{2, 3},
		{3, 2},
	}

	if canFinish(4, prerequisites) != false {
		t.Error("Expected dependent prerequisites not to be able to finish")
	}
}

package queuereconstructionbyheight

import (
	"testing"
)

func TestAlreadySortedOutput(t *testing.T) {
	list := [][]int{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
	}
	expected := [][]int{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
	}
	reconstruction := reconstructQueue(list)

	for i := 0; i < len(expected); i++ {
		inputItem := expected[i]
		outputItem := reconstruction[i]

		if inputItem[0] != outputItem[0] {
			t.Fatalf("Expected input item to be same for already sorted list: %v", reconstruction)
		}

		if inputItem[1] != outputItem[1] {
			t.Fatalf("Expected input item to be same for already sorted list: %v", reconstruction)
		}
	}
}

func TestResconstructs(t *testing.T) {
	list := [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}
	expected := [][]int{
		{5, 0},
		{7, 0},
		{5, 2},
		{6, 1},
		{4, 4},
		{7, 1},
	}
	reconstruction := reconstructQueue(list)

	for i := 0; i < len(expected); i++ {
		inputItem := expected[i]
		outputItem := reconstruction[i]

		if inputItem[0] != outputItem[0] {
			t.Fatalf("Expected input item to be same for already sorted list: %v", reconstruction)
		}

		if inputItem[1] != outputItem[1] {
			t.Fatalf("Expected input item to be same for already sorted list: %v", reconstruction)
		}
	}
}

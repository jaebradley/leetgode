package numberofislands

import "testing"

func TestEmpty2By2(t *testing.T) {
	grid := [][]byte{
		{byte('0'), byte('0')},
		{byte('0'), byte('0')},
	}

	islands := numIslands(grid)

	if islands != 0 {
		t.Errorf("Islands should be 0 and not %d", islands)
	}
}

func TestSingle2By2(t *testing.T) {
	grid := [][]byte{
		{byte('1'), byte('0')},
		{byte('0'), byte('0')},
	}

	islands := numIslands(grid)

	if islands != 1 {
		t.Errorf("Islands should be 1 and not %d", islands)
	}

}

func TestDouble2By2(t *testing.T) {
	grid := [][]byte{
		{byte('1'), byte('0')},
		{byte('0'), byte('1')},
	}

	islands := numIslands(grid)

	if islands != 2 {
		t.Errorf("Islands should be 2 and not %d", islands)
	}
}

func TestSingleIslandThreeElementsFor2By2(t *testing.T) {
	grid := [][]byte{
		{byte('1'), byte('0')},
		{byte('1'), byte('1')},
	}

	islands := numIslands(grid)

	if islands != 1 {
		t.Errorf("Islands should be 1 and not %d", islands)
	}
}

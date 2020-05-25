package verifypreorderserializationofabinarytree

import "testing"

func TestSingleLeftChildIsInvalid(t *testing.T) {
	if isValidSerialization("1,#") != false {
		t.Fatalf("Expected serialization validation to be false")
	}
}

func TestLeftAndRightChildIsValid(t *testing.T) {
	if isValidSerialization("1,2,#,#,3,#,#") != true {
		t.Fatalf("Expected serialization validation to be true")
	}
}

func TestCombination(t *testing.T) {
	if isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#") != true {
		t.Fatalf("Expected serialization validation to be true")
	}
}

func TestEmpty(t *testing.T) {
	if isValidSerialization("") != false {
		t.Fatalf("Expected serialization validation to be false")
	}
}

func TestNullRoot(t *testing.T) {
	if isValidSerialization("#") != true {
		t.Fatalf("Expected serialization validation to be true")
	}
}

func TestNullRootWithChildren(t *testing.T) {
	if isValidSerialization("#,#,3,5,#") != false {
		t.Fatalf("Expected serialization validation to be false")
	}
}

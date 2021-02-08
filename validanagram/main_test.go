package validanagram

import "testing"

func TestSingleCharacterString(t *testing.T) {
	if true != isAnagram("t", "t") {
		t.Fatalf("Expect t to be anagram of itself")
	}
}

func TestIdenticalTwoCharacterString(t *testing.T) {
	if true != isAnagram("ab", "ab") {
		t.Fatalf("Expect identical two character string to be anagram of itself")
	}
}

func TestDifferentTwoCharacterStrings(t *testing.T) {
	if false != isAnagram("aa", "bb") {
		t.Fatalf("Expect two different strings are not anagrams")
	}
}

func TestBackwardsCharacterString(t *testing.T) {
	if true != isAnagram("ab", "ba") {
		t.Fatalf("Expect backwards string to be anagram")
	}
}

func TestMultipleOfSameCharacter(t *testing.T) {
	if true != isAnagram("abb", "bba") {
		t.Fatalf("Expect multiple of same character to be anagram")
	}
}

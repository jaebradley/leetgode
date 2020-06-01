package wordbreak

import "testing"

func TestSingleWordMatchesDictionary(t *testing.T) {
	if wordBreak("foobar", []string{"foobar"}) != true {
		t.Errorf("Expected foobar to match dictionary of [foobar]")
	}
}

func TestSingleCharacterMatchesDictionaryWithSingleCharacter(t *testing.T) {
	if wordBreak("a", []string{"a"}) != true {
		t.Errorf("Expected a to match dictionary of [a]")
	}
}

func TestSingleCharacterNotMatchingDictionaryWithSingleCharacter(t *testing.T) {
	if wordBreak("a", []string{"b"}) != false {
		t.Errorf("Expected a not to match dictionary of [b]")
	}
}

func TestMultipleSubstringMatchesWithoutMatchingEntireString(t *testing.T) {
	if wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) != false {
		t.Errorf("Expected catsanddogs not to match a dictionary of ['cats', 'dog', 'sand', 'and', 'cat']")
	}
}

package wordbreak

/*
https://leetcode.com/problems/word-break/

Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Approach:

* Keep boolean array that is length of string
* Set the first index of boolean array to true
* Iterate through string starting from first index - this will be start index of all substrings
* Iterate through all indices greater than first index - this will be end index of all substring
* Only consider substrings where the start index was considered valid - in other words, the words used to get to that substring all exist in the dictionary
* If substring exists in dictionary then mark the index for that word as valid (i.e. mark it as true in boolean array)
* After iterating through all possible combinations, if last index is valid then the word can be broken up using the dictionary
*/

func wordBreak(s string, wordDict []string) bool {
	validIndices := make([]bool, len(s)+1)
	validIndices[0] = true

	for endIndex := 1; endIndex <= len(s); endIndex++ {
		for startIndex := 0; startIndex < endIndex; startIndex++ {
			if validIndices[startIndex] {
				substring := s[startIndex:endIndex]
				for _, word := range wordDict {
					if substring == word {
						validIndices[endIndex] = true
						break
					}
				}
			}
		}
	}

	return validIndices[len(s)]
}

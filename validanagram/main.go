package validanagram

func isAnagram(s string, t string) bool {
	characterCounts := make(map[rune]int)

	for _, c := range s {
		characterCounts[c]++
	}

	for _, c := range t {
		characterCounts[c]--
	}

	for _, v := range characterCounts {
		if 0 != v {
			return false
		}
	}

	return true
}

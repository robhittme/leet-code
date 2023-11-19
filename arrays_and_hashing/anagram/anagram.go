package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := make(map[rune]int)
	for _, c := range s {
		if hash[c] > 0 {
			hash[c]++
		} else {
			hash[c] = 1
		}
	}
	for _, b := range t {
		if hash[b] > 0 {
			hash[b]--
		} else {
			return false
		}
	}
	return true
}

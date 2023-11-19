package main

func isAnagram1(s string, t string) bool {
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
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m [26]int
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

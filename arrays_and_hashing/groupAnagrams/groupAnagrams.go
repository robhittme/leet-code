/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]


Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/

package main

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

func groupAnagrams(strs []string) [][]string {
	var hash [][]string
	for i, n := range strs {
		var found bool
		for j, v := range hash {
			if len(v) > 0 && isAnagram(n, v[0]) {
				hash[j] = append(hash[j], n)
				found = true
				break
			}
		}
		if !found {
			hash = append(hash, []string{strs[i]})
		}
	}
	return hash
}

func groupAnagrams1(strs []string) [][]string {
	freqs := make(map[[26]byte][]string, len(strs))

	for _, s := range strs {
		g := [26]byte{}
		for _, c := range s {
			g[c-'a']++
		}
		freqs[g] = append(freqs[g], s)
	}

	result := [][]string{}

	for _, v := range freqs {
		result = append(result, v)
	}
	return result
}

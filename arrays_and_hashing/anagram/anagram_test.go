package main

import "testing"

var testCases = []struct {
	input1   string
	input2   string
	expected bool
}{
	{"listen", "silent", true},
	{"heart", "earth", true},
	{"ab", "a", false},
	{"top", "stop", false},
}

func TestIsAnagram1(t *testing.T) {
	for _, tc := range testCases {
		result := isAnagram1(tc.input1, tc.input2)
		if result != tc.expected {
			t.Errorf("isAnagram(%v, %v) = %v; expected %v", tc.input1, tc.input2, result, tc.expected)
		}
	}

}

func TestIsAnagram(t *testing.T) {
	for _, tc := range testCases {
		result := isAnagram(tc.input1, tc.input2)
		if result != tc.expected {
			t.Errorf("isAnagram(%v, %v) = %v; expected %v", tc.input1, tc.input2, result, tc.expected)
		}
	}

}

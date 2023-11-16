package main

import "testing"

var testCases = []struct {
	input    []int
	expected bool
}{
	{[]int{1, 2, 3, 1}, true},    // Duplicates
	{[]int{1, 2, 3, 4}, false},   // No duplicates
	{[]int{5, 5, 5, 5, 5}, true}, // All elements are the same (duplicates)
	{[]int{}, false},             // Empty slice
}

func TestContainsDuplicates(t *testing.T) {
	for _, tc := range testCases {
		result := containsDuplicates(tc.input)
		if result != tc.expected {
			t.Errorf("containsDuplicates(%v) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}

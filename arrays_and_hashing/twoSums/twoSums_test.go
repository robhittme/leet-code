package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	arr      []int
	target   int
	expected []int
}{
	{[]int{1, 2, 3, 1}, 5, []int{1, 2}},
	{[]int{1, 2, 3, 4}, 7, []int{2, 3}},
}

func TestTwoSums(t *testing.T) {
	for i, tc := range testCases {
		result := twoSums(tc.arr, tc.target)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Test case %d failed. Expected %v, but got %v for arr=%v, target=%d", i+1, tc.expected, result, tc.arr, tc.target)
		}
	}
}

package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	input    []int
	target   int
	expected []int
}{
	{[]int{1, 1, 2, 2, 3}, 2, []int{1, 2}},
	{[]int{1}, 1, []int{1}},
	{[]int{1, 2}, 2, []int{1, 2}},
}

func TestTopKFrequent(t *testing.T) {
	for _, tc := range testCases {
		result := topKFrequent(tc.input, tc.target)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("topKFrequent(%v) expected %v, but got %v", tc.input, tc.expected, result)
		}
	}
}


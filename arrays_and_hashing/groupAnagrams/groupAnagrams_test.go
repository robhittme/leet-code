package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	input    []string
	expected [][]string
}{
	{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
	{[]string{""}, [][]string{{""}}},
	{[]string{"a"}, [][]string{{"a"}}},
	{[]string{"abc", "def"}, [][]string{{"abc"}, {"def"}}},
	{[]string{"listen", "silent", "hello", "world"}, [][]string{{"listen", "silent"}, {"hello"}, {"world"}}},
}

func TestGroupAnagrams(t *testing.T) {
	for _, tc := range testCases {
		result := groupAnagrams(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("groupAnagrams(%v) expected %v, but got %v", tc.input, tc.expected, result)
		}
	}
}

func TestGroupAnagrams1(t *testing.T) {
	for _, tc := range testCases {
		result := groupAnagrams1(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("groupAnagrams(%v) expected %v, but got %v", tc.input, tc.expected, result)
		}
	}
}

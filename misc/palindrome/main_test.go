package main

import (
  "testing"
)


func TestPalindrome(t *testing.T) {
  cases := []struct {
    input string
    expected bool
  }{
    {"racecar", true},
    {"asdfdsa", true},
    {"A man, a plan, a canal: Panama", true},
    {" ", true},
    {"race a car", false},
  }
 
  for _, c := range cases {
    v := isPalindrome(c.input)    
    if v != c.expected {
      t.Errorf("FAILED: was expecting %t but got %t", c.expected, v) 
    }
  }
}

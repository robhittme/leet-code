package main

import (
    "fmt"
    "regexp"
    "strings"
)

func isPalindrome(s string) bool {
  reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	result := reg.ReplaceAllString(s, "")
  output := []rune{}
  for _, v := range result {
    output = append([]rune{v}, output...)
  }
  fmt.Println(string(output), result)
  return strings.ToLower(string(output)) == strings.ToLower(result)
}

func main() {
    s := "A man, a plan, a canal: Panama"
    fmt.Println(isPalindrome(s))
}

package main

import (
    "fmt"
)

func romanToInt(s string) int {
    sum := 0
    var prev rune 
    mapping := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    for _, val := range s {
        sum += mapping[val]
        if mapping[prev] < mapping[val] {
            sum -= mapping[prev] * 2
        }
        prev = val
    }
    return sum
}

func main() {

    s := "MCMXCIV"
    fmt.Println(romanToInt(s))
}

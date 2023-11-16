package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	plus := false
    for i := len(digits) - 1; i >= 0; i-- {
		if plus {
			return digits
		}
		if digits[i] != 9 {
			digits[i] = digits[i] + 1  
			plus = true
		} else {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		}

    }
	return digits 
}
func main() {
	data := []int{9}
	fmt.Println(plusOne(data))

}

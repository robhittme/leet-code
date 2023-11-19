package main

import "fmt"

func productOfArray(nums []int) []int {
	res := make([]int, len(nums))
	prefix := 1
	for i := range nums {
		res[i] = prefix
		prefix = prefix * nums[i]
		fmt.Println(res)
	}
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * postfix
		postfix = postfix * nums[i]
		fmt.Println(res)
	}

	return res
}

func main() {
	arr := []int{1, 2, 3, 4}
	val := productOfArray(arr)
	fmt.Println(val)
}

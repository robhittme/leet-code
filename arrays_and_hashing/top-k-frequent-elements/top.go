package main

func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)
	for _, n := range nums {
		dict[n]++
	}

	return []int{}
}

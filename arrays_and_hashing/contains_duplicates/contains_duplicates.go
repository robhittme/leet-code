package main

func containsDuplicates(val []int) bool {
	hash := make(map[int]int)
	for _, n := range val {
		if hash[n] == 1 {
			return true
		}
		hash[n] = 1
	}
	return false
}

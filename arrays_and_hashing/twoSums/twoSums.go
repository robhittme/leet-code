package main

//	func twoSums(arr []int, target int) []int {
//		hash := make(map[int]int)
//		for i, n := range arr {
//			remainder := target - n
//			if r, ok := hash[remainder]; ok {
//				return []int{r, i}
//			}
//			hash[n] = i
//		}
//		return []int{}
//	}
func twoSums(arr []int, target int) []int {
	hash := make(map[int]int)
	for i, n := range arr {
		remainder := target - n
		if r, ok := hash[remainder]; ok {
			return []int{r, i}
		}
		hash[n] = i
	}
	return []int{}
}

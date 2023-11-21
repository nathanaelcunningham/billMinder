package main

func numList(start, end int) []int {
	var nums []int
	for i := start; i <= end; i++ {
		nums = append(nums, i)
	}
	return nums
}

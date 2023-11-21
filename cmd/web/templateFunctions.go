package main

import (
	"fmt"
	"time"
)

func numList(start, end int) []int {
	var nums []int
	for i := start; i <= end; i++ {
		nums = append(nums, i)
	}
	return nums
}

func getDate(dayOfMonth int64) string {
	curr := time.Now()
	date := time.Date(curr.Year(), curr.Month(), int(dayOfMonth), 0, 0, 0, 0, time.UTC)
	return fmt.Sprintf("%s %d", date.Format("Monday"), dayOfMonth)
}

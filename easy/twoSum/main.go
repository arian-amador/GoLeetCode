package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	n := []int{15, 7, 2, 11}
	t := 91

	go func() {
		defer wg.Done()
		fmt.Printf("Non-Sorted Two Sums Results: %+v = %d\n", bruteTwoSum(n, t), t)
	}()

	go func() {
		defer wg.Done()
		fmt.Printf("Sorted Two Sums Results: %+v = %d\n", sortedTwoSum(n, t), t)
	}()

	go func() {
		defer wg.Done()
		fmt.Printf("One loop hash lookup Two Sums Results: %+v = %d\n", onePassHash(n, t), t)
	}()

	wg.Wait()
}

func onePassHash(nums []int, target int) []int {
	l := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		c := target - nums[i]
		if val, ok := l[c]; ok {
			return []int{val, i}
		}
		l[nums[i]] = i

	}
	return []int{}
}

func bruteTwoSum(nums []int, target int) []int {
	res := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = []int{i, j}
				return res
			}
		}
	}
	return res
}

func sortedTwoSum(nums []int, target int) []int {
	sort.Ints(nums)

	l := 0
	r := len(nums) - 1
	res := []int{}

	for i := 0; i < len(nums); i++ {
		temp := nums[l] + nums[r]

		if temp == target {
			res = []int{l, r}
			break
		}

		if temp > target {
			r--
		} else {
			l++
		}
	}

	return res
}

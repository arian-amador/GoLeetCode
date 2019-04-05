package twoSum

import (
	"sort"
)

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

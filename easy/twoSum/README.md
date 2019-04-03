# EASY - Two Sum [![Build Status](https://api.travis-ci.org/arian-amador/GoLeetCode.svg)](https://travis-ci.org/arian-amador/GoLeetCode)

### Problem

Given an array of integers, return indices of the two numbers such that they add up to a specific target. You may assume that each input would have exactly one solution, and you may not use the same element twice.

##### Example

```Go
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]
```

### Solutions

##### Brute Force

The simplest solution to jump for is the nested for loop solution.
It's the easiest to reason about but also ends up being the slowest by quite a bit.

```Go
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
```

##### First Sorted

This attempt we first sort the original list of integers and itterate from both ends.
Keeping track on the the head and tail of the slice we're able to adjust them accordingly.
If the sum is less than the target we move the left side otherwise we move the right side.

```Go
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
```

##### Single Pass Lookup Table

The most efficient way of solving this exercise is to first create a has to be our lookup table.
This hash will store the difference of the desired _sum_ and the current integer from the slice.
If the difference equals a previously stored value we know the current integer and the previously stored one would equal the desired target integer.
Otherwise we store the current integer as the key with the index as the value.

```Go
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
```

---

#### Hope you find this useful!

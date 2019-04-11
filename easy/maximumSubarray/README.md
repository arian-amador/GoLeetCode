# EASY - Maximum Subarray [![Build Status](https://api.travis-ci.org/arian-amador/GoLeetCode.svg)](https://travis-ci.org/arian-amador/GoLeetCode)

### Problem

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

##### Example

```Go
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

### Solutions

The key to this solution is understanding Kadane's algorithm.
Kadane's algorithm assumes that you can update the current max element by comparing the
previous max subarray with the current element.
Taking the `max` from those options gives you the largest subarray.
The final bit is continuously updating a max counter.

```Go
for i := 1; i <= len(nums)-1; i++ {
  if nums[i] < curr+nums[i] {
    curr = curr + nums[i]
  } else {
    curr = nums[i]
  }

  if curr > max {
    max = curr
  }
}
```

---

#### Hope you find this useful!

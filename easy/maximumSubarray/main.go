package maxSubarray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	curr := max

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

	return max
}

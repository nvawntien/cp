func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	var dp func(i, total int) int
	dp = func(i, total int) int {
		if i == n {
			if total == target {
				return 1
			}
			return 0
		}

		return dp(i+1, total+nums[i]) + dp(i+1, total-nums[i])
	}

	return dp(0, 0)
}
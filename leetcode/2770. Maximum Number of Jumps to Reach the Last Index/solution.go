func maximumJumps(nums []int, target int) int {
    n := len(nums)
    dp := make([]int, n)

    for i := range dp {
        dp[i] = -1
    }

    dp[0] = 0
    
    for j := 0; j < n; j++ {
        for i := 0; i < j; i++ {
            if dp[i] != -1 && nums[i] - target <= nums[j] && nums[j] <= nums[i] + target {
                dp[j] = max(dp[j], dp[i]+1)
            }
        }
    }

    return dp[n-1]
}
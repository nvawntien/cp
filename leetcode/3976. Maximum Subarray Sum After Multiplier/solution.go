func maxSubarraySum(nums []int, k int) int64 {
    n := len(nums)
    c := int64(nums[0])
    m := int64(nums[0] * k)
    d := int64(nums[0] / k)
    
    dp := [5]int64 {c, m, m, d, d}
    ans := max(max(c, d), m)
    
    for i := 1; i < n; i++ {
        c = int64(nums[i])
        m = int64(nums[i] * k)
        d = int64(nums[i] / k)

        next := [5]int64{}

        next[0] = max(c, dp[0] + c)
        next[1] = max(m, max(dp[1] + m, dp[0] + m))
        next[2] = max(c, max(dp[1] + c, dp[2] + c))
        next[3] = max(d, max(dp[3] + d, dp[0] + d))
        next[4] = max(c, max(dp[3] + c, dp[4] + c))

        dp = next
        for i := range next {
            ans = max(ans, next[i])
        }
    }

    return ans
}
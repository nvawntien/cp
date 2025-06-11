func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	return min(dp[n-1], dp[n-2])
}
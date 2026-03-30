func countBits(n int) []int {
	dp := make([]int, n+1)
	powerOfTwo := []int{}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			dp[i] = 1
			powerOfTwo = append(powerOfTwo, i)
			continue
		}
		a := powerOfTwo[len(powerOfTwo)-1]
		b := i - a
		dp[i] = dp[a] + dp[b]
	}

	return dp
}
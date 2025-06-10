func isSubsequence(s string, t string) bool {
	n := len(s)
	m := len(t)

	s = " " + s
	t = " " + t

	dp := make([][]bool, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	for i := 0; i <= m; i++ {
		dp[0][i] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i][j-1]
			if j >= i {
				temp := false
				if s[i] == t[j] {
					temp = true
				}

				dp[i][j] = dp[i][j] || (dp[i-1][j-1] && temp)
			}
		}
	}

	return dp[n][m]
}
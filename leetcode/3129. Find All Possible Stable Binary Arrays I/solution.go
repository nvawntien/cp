func numberOfStableArrays(zero int, one int, limit int) int {
    // dp[zero+1][one+1][k={0,1}]
    MOD := 1_000_000_007
    dp := make([][][]int, zero+1)

    for i := range dp {
        dp[i] = make([][]int, one+1)
        for j := range dp[i] {
            dp[i][j] = make([]int, 2)
        }
    }

    for i := 0; i <= min(zero, limit); i++ {
        dp[i][0][0] = 1
    }

    for j := 0; j <= min(one, limit); j++ {
        dp[0][j][1] = 1
    }

    for i := 1; i <= zero; i++ {
        for j := 1; j <= one; j++ {
            dp[i][j][0] = (dp[i-1][j][0] + dp[i-1][j][1]) % MOD
            if i > limit {
                dp[i][j][0] = (dp[i][j][0] - dp[i-limit-1][j][1] + MOD) % MOD
            }

            dp[i][j][1] = (dp[i][j-1][0] + dp[i][j-1][1]) % MOD
            if j > limit {
                dp[i][j][1] = (dp[i][j][1] - dp[i][j-limit-1][0] + MOD) % MOD
            }
        }
    }

    return (dp[zero][one][0] + dp[zero][one][1]) % MOD
}
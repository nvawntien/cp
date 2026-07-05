func zigZagArrays(n int, l int, r int) int {
    m := r - l + 1    
    if n == 1 {
        return m
    }

    dp := make([][][2]int, n+1)
    for i := range dp {
        dp[i] = make([][2]int, m+2)
    }

    mod := 1_000_000_007

    for j := 1; j <= m; j++ {
        dp[1][j][0] = 1
        dp[1][j][1] = 1
    }

    for i := 2; i <= n; i++ {
        pre := 0
        suf := 0
        for j := 1; j <= m; j++ {
            pre = (pre + dp[i-1][j-1][1]) % mod
            dp[i][j][0] = (dp[i][j][0] + pre) % mod
        }

        for j := m; j >= 1; j-- {
            suf = (suf + dp[i-1][j+1][0]) % mod
            dp[i][j][1] = (dp[i][j][1] + suf) % mod

        }
    }

    ans := 0
    for j := 1; j <= m; j++ {
        ans = (ans + dp[n][j][0]) % mod
        ans = (ans + dp[n][j][1]) % mod
    }

    return ans
}
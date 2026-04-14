const INF = int64(1e12)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
    sort.Ints(robot)
    sort.Slice(factory, func (i, j int) bool {
        return factory[i][0] < factory[j][0]
    })

    n := len(robot)
    m := len(factory)
    dp := make([][]int64, n+1)
    pre := make([][]int64, n+1)

    abs := func (x int) int {
        mask := x >> 63
        return (x ^ mask) - mask
    }

    for i := range dp {
        dp[i] = make([]int64, m+1)
        pre[i] = make([]int64, m+1)
        for j := range dp[i] {
            dp[i][j] = INF
        }
    }

    for j := 1; j <= m; j++ {
        for i := 1; i <= n; i++ {
            pre[i][j] = pre[i-1][j] + int64(abs(factory[j-1][0] - robot[i-1]))
        }
    }
    
    for j := 0; j <= m; j++ {
        dp[0][j] = 0
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            for k := 0; k <= factory[j-1][1]; k++ {
                if i-k >= 0 {
                    dp[i][j] = min(dp[i][j], dp[i-k][j-1] + pre[i][j] - pre[i-k][j])
                }
            }
        }
    }

    var ans = INF

    for j := range dp[n] {
        ans = min(ans, dp[n][j])
    }

    return ans
}

// 0 4 6

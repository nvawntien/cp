func maxProductPath(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])

    dp := make([][][]int, n)

    for i := range dp {
        dp[i] = make([][]int, m)
        for j := range dp[i] {
            dp[i][j] = make([]int, 2)
        }
    } 

    dp[0][0][0] = grid[0][0]
    dp[0][0][1] = grid[0][0]

    for i := 1; i < n; i++ {
        if grid[i][0] >= 0 {
            dp[i][0][0] = dp[i-1][0][0] * grid[i][0]
            dp[i][0][1] = dp[i-1][0][1] * grid[i][0]
        } else {
            dp[i][0][1] = dp[i-1][0][0] * grid[i][0]
            dp[i][0][0] = dp[i-1][0][1] * grid[i][0]
        }
    }  

    
    for j := 1; j < m; j++ {
        if grid[0][j] >= 0 {
            dp[0][j][0] = dp[0][j-1][0] * grid[0][j]
            dp[0][j][1] = dp[0][j-1][1] * grid[0][j]
        } else {
            dp[0][j][0] = dp[0][j-1][1] * grid[0][j]
            dp[0][j][1] = dp[0][j-1][0] * grid[0][j]
        }
    }

    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if grid[i][j] >= 0 {
                dp[i][j][0] = max(dp[i-1][j][0], dp[i][j-1][0]) * grid[i][j]
                dp[i][j][1] = min(dp[i-1][j][1], dp[i][j-1][1]) * grid[i][j]
            } else {
                dp[i][j][0] = min(dp[i-1][j][1], dp[i][j-1][1]) * grid[i][j]
                dp[i][j][1] = max(dp[i-1][j][0], dp[i][j-1][0]) * grid[i][j]
            }
        }
    }

    if dp[n-1][m-1][0] < 0 {
        return -1
    }

    return dp[n-1][m-1][0] % 1_000_000_007
}
func maxPathScore(grid [][]int, k int) int {
    n := len(grid)
    m := len(grid[0])
    dp := make([][][]int, n)

    for i := range dp {
        dp[i] = make([][]int, m)
        for j := range dp[i] {
            dp[i][j] = make([]int, k+1)
            for c := range dp[i][j] {
                dp[i][j][c] = -1
            }
        }
    }

    costs := map[int]int{
        0: 0,
        1: 1,
        2: 1,
    }
    
    startCost := costs[grid[0][0]]
    
    if startCost <= k {
        dp[0][0][startCost] = grid[0][0]
    }


    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if i == 0 && j == 0 {
                continue
            }

            cost := costs[grid[i][j]]
            score := grid[i][j]
            for c := cost; c <= k; c++ {
                p := c - cost
              
                if i > 0 && dp[i-1][j][p] != -1 {
                    dp[i][j][c] = max(dp[i][j][c], dp[i-1][j][p] + score)
                }

                if j > 0 && dp[i][j-1][p] != -1 {
                    dp[i][j][c] = max(dp[i][j][c], dp[i][j-1][p] + score)
                }
            }
        }
    }

    ans := -1
    for _, score := range dp[n-1][m-1] {
        ans = max(ans, score)
    }

    return ans
}   
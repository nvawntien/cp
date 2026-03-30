func countSubmatrices(grid [][]int, k int) int {
    n := len(grid)
    m := len(grid[0])

    pre := make([][]int, n)

    for i := range pre {
        pre[i] = make([]int, m)
    }

    pre[0][0] = grid[0][0]

    for i := 1; i < n; i++ {
        pre[i][0] = pre[i-1][0] + grid[i][0]
    }

    for j := 1; j < m; j++ {
        pre[0][j] = pre[0][j-1] + grid[0][j]
    }

    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + grid[i][j]
        }
    }

    cnt := 0

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if pre[i][j] <= k {
                cnt++
            }
        }
    }

    return cnt
}
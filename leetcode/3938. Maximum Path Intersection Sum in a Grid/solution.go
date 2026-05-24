func maxScore(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])

    ans := math.MinInt

    if n == 1 || m == 1 {
        sum := 0
        for i := 0; i < n; i++ {
            for j := 0; j < m; j++ {
                sum += grid[i][j]
            }
        }
        return sum
    }
    
    for i := 0; i < n; i++ {
        curr := grid[i][0]
        for j := 1; j < m; j++ {
            val := grid[i][j]
            if curr + val > ans {
                ans = curr + val
            }

            curr = max(val, curr + val)
        }    
    }

    for j := 0; j < m; j++ {
        curr := grid[0][j]
        for i := 1; i < n; i++ {
            val := grid[i][j]
            if curr + val > ans {
                ans = curr + val
            }

            curr = max(val, curr + val)
        }    
    }


    for i := 1; i < n-1; i++ {
        for j := 1; j < m-1; j++ {
            if grid[i][j] > ans {
                ans = grid[i][j]
            }
        }
    }
    return ans
}
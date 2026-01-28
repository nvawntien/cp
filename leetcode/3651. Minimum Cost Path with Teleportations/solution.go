type Point struct {
    r, c int
}

func minCost(grid [][]int, k int) int {
    m := len(grid)
    n := len(grid[0])
    dp := make([][][]int, k+1)

    for i := range dp {
        dp[i] = make([][]int, m)
        for r := range dp[i] {
            dp[i][r] = make([]int, n)
            for c := range dp[i][r] {
                dp[i][r][c] = math.MaxInt64
            }
        }
    }

    dp[0][0][0] = 0
    groups := make(map[int][]Point)

    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            val := grid[r][c]
            groups[val] = append(groups[val], Point{r, c})
        }
    }

    var sortedVals []int 
    for val := range groups {
        sortedVals = append(sortedVals, val)
    }

    sort.Slice(sortedVals, func (i, j int) bool {
        return sortedVals[i] > sortedVals[j]
    })

    ans := math.MaxInt64 

    for t := 0; t <= k; t++ {
        for r := 0; r < m; r++ {
            for c := 0; c < n; c++ {
                if r == 0 && c == 0 {
                    continue
                }

                if r-1 >= 0 && dp[t][r-1][c] != math.MaxInt64 {
                    dp[t][r][c] = min(dp[t][r][c], dp[t][r-1][c] + grid[r][c])
                }    

                if c-1 >= 0 && dp[t][r][c-1] != math.MaxInt64 {
                    dp[t][r][c] = min(dp[t][r][c], dp[t][r][c-1] + grid[r][c])
                }
            }
        }

        ans = min(ans, dp[t][m-1][n-1])

        if t < k {
            minCostFound := math.MaxInt64
            
            for _, val := range sortedVals {
                points := groups[val]
                curr := math.MaxInt64
                for _, p := range points {
                    curr = min(curr, dp[t][p.r][p.c])
                }

                minCostFound = min(minCostFound, curr)
                if minCostFound != math.MaxInt64 {
                    for _, p := range points {
                        dp[t+1][p.r][p.c] = min(dp[t+1][p.r][p.c], minCostFound)
                    }
                }
            }
        }
    }

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
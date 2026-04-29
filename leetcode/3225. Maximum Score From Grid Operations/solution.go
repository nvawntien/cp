func maximumScore(grid [][]int) int64 {
    n := len(grid)
    pre := make([][]int64, n)

    for i := range pre {
        pre[i] = make([]int64, n+1)
        for j := 0; j < n; j++ {
            pre[i][j+1] = pre[i][j] + int64(grid[j][i])
        }
    }

    prevPick := make([]int64, n+1)
    prevSkip := make([]int64, n+1)

    for j := 1; j < n; j++ {
        currPick := make([]int64, n+1)
        currSkip := make([]int64, n+1)

        for curr := 0; curr <= n; curr++ {
            for prev := 0; prev <= n; prev++ {
                if prev < curr {
                    score := pre[j-1][curr] - pre[j-1][prev]
                    currPick[curr] = max(currPick[curr], prevSkip[prev] + score)
                    currSkip[curr] = max(currSkip[curr], prevSkip[prev] + score)
                } else {
                    score := pre[j][prev] - pre[j][curr]
                    currPick[curr] = max(currPick[curr], prevPick[prev] + score)
                    currSkip[curr] = max(currSkip[curr], prevPick[prev])
                }
            }
        }

        prevPick = currPick
        prevSkip = currSkip
    }

    var ans int64 = 0
    for _, x := range prevPick {
        ans = max(ans, x)
    }

    return ans
}
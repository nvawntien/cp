func maxSideLength(mat [][]int, threshold int) int {
    n := len(mat)
    m := len(mat[0])

    pre := make([][]int, n+1)
    for i := range pre {
        pre[i] = make([]int, m+1)
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            pre[i][j] = mat[i-1][j-1] + pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1]
        }
    }

    low, high := 1, min(n, m)
    var ans = 0

    for low <= high {
        mid := (low + high) >> 1
        if hasValidSquare(mid, threshold, n, m, pre) {
            ans = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return ans
}

func hasValidSquare(side, threshold, n, m int, pre [][]int) bool {
    for r := side; r <= n; r++ {
        for c := side; c <= m; c++ {
            sum := pre[r][c] - pre[r-side][c] - pre[r][c-side] + pre[r-side][c-side];
            if sum <= threshold {
                return true
            }
        }
    }

    return false
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
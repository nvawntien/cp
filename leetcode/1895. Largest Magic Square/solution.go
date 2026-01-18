func largestMagicSquare(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])

    rowSum := make([][]int, n)
    
    for i := 0; i < n; i++ {
        rowSum[i] = make([]int, m+1)
        for j := 0; j < m; j++ {
            rowSum[i][j+1] = rowSum[i][j] + grid[i][j]
        }
    }

    colSum := make([][]int, n+1)

    for i := range colSum {
        colSum[i] = make([]int, m)
    }

    for j := 0; j < m; j++ {
        for i := 0; i < n; i++ {
            colSum[i+1][j] = colSum[i][j] + grid[i][j]
        } 
    }

    for k := min(n, m); k > 1; k-- {
        for r := 0; r <= n-k; r++ {
            for c := 0; c <= m-k; c++ {
                if isMagic(grid, r, c, k, rowSum, colSum) {
                    return k
                }
            }
        }
    }

    return 1
}

func isMagic(grid [][]int, r, c, k int, rowSum, colSum [][]int) bool {
    target := rowSum[r][c+k] - rowSum[r][c]
    
    for i := r+1; i < r+k; i++ {
        if rowSum[i][c+k] - rowSum[i][c] != target {
            return false
        }
    }

    for i := c; i < c+k; i++ {
        if colSum[r+k][i] - colSum[r][i] != target {
            return false
        }
    }

    diag1, diag2 := 0, 0
    
    for i := 0; i < k; i++ {
        diag1 += grid[r+i][c+i]
        diag2 += grid[r+i][c+k-i-1]
    }

    if diag1 != target || diag2 != target{
        return false
    }

    return true;
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}
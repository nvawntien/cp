func constructProductMatrix(grid [][]int) [][]int {
    n := len(grid)
    m := len(grid[0])
    mod := 12345    

    pre := make([]int, n*m)
    suf := make([]int, n*m)

    mul := 1
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            pre[i * m + j] = (mul * grid[i][j]) % mod
            mul = pre[i * m + j]
        }
    }

    // 1 2
    // 3 4
    // 1 2 3 4

    mul = 1
    for i := n-1; i >= 0; i-- {
        for j := m-1; j >= 0; j-- {
            suf[i * m + j] = (mul * grid[i][j]) % mod
            mul = suf[i * m + j]
        }
    }
    prod := make([][]int, n)
    for i := range prod {
        prod[i] = make([]int, m)
    }
    
    for i := range prod {
        for j := range prod[i] {
            var ans int
            if i == 0 && j == 0 {
                ans = suf[1] % mod
            } else if i == n-1 && j == m-1 {
                ans = pre[n*m-2] % mod
            } else {
                ans = (pre[i*m+j-1] * suf[i*m+j+1]) % mod
            }
            prod[i][j] = ans
        }
    }

    return prod
}
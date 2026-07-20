func shiftGrid(grid [][]int, k int) [][]int {
    m := len(grid)
    n := len(grid[0])

    new := make([][]int, m)
    for i := range new {
        new[i] = make([]int, n)
    }
    // (i * n + j + k) % (n * m -1) - 1
    // (a - (a / m)) / n = j
    // i = a / m
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            idx := (i * n + j + k) % (n * m)
            ni := idx/n
            nj := idx%n
            new[ni][nj] = grid[i][j]
        }
    }

    return new
}
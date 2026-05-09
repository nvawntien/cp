func rotateGrid(grid [][]int, k int) [][]int {
    n := len(grid)
    m := len(grid[0])
    layer := min(n, m) >> 1

    for l := 0; l < layer; l++ {
        var arr []int
        
        for j := l; j < m-l-1; j++ {
            arr = append(arr, grid[l][j])
        }

        for i := l; i < n-l-1; i++ {
            arr = append(arr, grid[i][m-l-1])
        }

        for j := m-l-1; j > l; j-- {
            arr = append(arr, grid[n-l-1][j])
        }

        for i := n-l-1; i > l; i-- {
            arr = append(arr, grid[i][l])
        }

        length := len(arr)
        shift := k % length
        idx := 0
        
        for j := l; j < m-l-1; j++ {
            grid[l][j] = arr[(idx + shift) % length]
            idx++
        }

        for i := l; i < n-l-1; i++ {
            grid[i][m-l-1] = arr[(idx + shift) % length]
            idx++
        }

        for j := m-l-1; j > l; j-- {
            grid[n-l-1][j] = arr[(idx + shift) % length]
            idx++
        }

        for i := n-l-1; i > l; i-- {
            grid[i][l] = arr[(idx + shift) % length]
            idx++
        }
    }

    return grid
}
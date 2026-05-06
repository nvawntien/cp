func rotateTheBox(boxGrid [][]byte) [][]byte {
    n := len(boxGrid)
    m := len(boxGrid[0])

    rotate := make([][]byte, m)
    for i := range rotate {
        rotate[i] = make([]byte, n)
    }

    for i := 0; i < n; i++ {
        swap := m-1
        for j := m-1; j >= 0; j-- {
            if boxGrid[i][j] == '*' {
                swap = j-1
            } else if boxGrid[i][j] == '#' {
                boxGrid[i][j], boxGrid[i][swap] = boxGrid[i][swap], boxGrid[i][j]
                swap--
            }
        }

        for j := 0; j < m; j++ {
            rotate[j][n-i-1] = boxGrid[i][j]
        }
    }

    return rotate
}
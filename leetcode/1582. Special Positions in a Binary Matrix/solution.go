func numSpecial(mat [][]int) int {
    spec := 0
    n := len(mat)
    m := len(mat[0])

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if mat[i][j] == 1 && check(mat, i, j, n, m){
                spec++
            }
        }
    }

    return spec
}

func check(mat [][]int, r, c, n, m int) bool {
    for i := 0; i < m; i++ {
        if i != c && mat[r][i] == 1 {
            return false
        }
    }

    for i := 0; i < n; i++ {
        if i != r && mat[i][c] == 1 {
            return false
        }
    }

    return true
}
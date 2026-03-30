func numberOfSubmatrices(grid [][]byte) int {
    n := len(grid)
    m := len(grid[0])
    
    preX := make([][]int, n)
    preY := make([][]int, n)
    preDot := make([][]int, n)

    for i := 0; i < n; i++ {
        preX[i] = make([]int, m)
        preY[i] = make([]int, m)
        preDot[i] = make([]int, m)
    }
    check := func (c byte, i, j int) int {
        if grid[i][j] == c {
            return 1
        }
        return 0
    }

    preX[0][0] = check('X', 0, 0)
    preY[0][0] = check('Y', 0, 0)
    preDot[0][0] = check('.', 0, 0)

    for i := 1; i < n; i++ {
        preX[i][0] = preX[i-1][0] + check('X', i, 0)
        preY[i][0] = preY[i-1][0] + check('Y', i, 0)
        preDot[i][0] = preDot[i-1][0] + check('.', i, 0)        
    }

    for j := 1; j < m; j++ {
        preX[0][j] = preX[0][j-1] + check('X', 0, j)
        preY[0][j] = preY[0][j-1] + check('Y', 0, j)
        preDot[0][j] = preDot[0][j-1] + check('.', 0, j)        
    }

    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            preX[i][j] = preX[i-1][j] + preX[i][j-1] - preX[i-1][j-1] + check('X', i, j)
            preY[i][j] = preY[i-1][j] + preY[i][j-1] - preY[i-1][j-1] + check('Y', i, j)
            preDot[i][j] = preDot[i-1][j] + preDot[i][j-1] - preDot[i-1][j-1] + check('.', i, j)
        }
    }

    cnt := 0

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if preX[i][j] > 0 && preX[i][j] == preY[i][j] {
                cnt++
            }
        }
    }

    return cnt
}


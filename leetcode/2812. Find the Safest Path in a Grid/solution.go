func maximumSafenessFactor(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])
    d := make([][]int, n)
    
    for i := range d {
        d[i] = make([]int, m)
        for j := range d[i] {
            d[i][j] = math.MaxInt
        }
    }

    dir := [5]int {0, 1, 0, -1, 0}  

    queue := [][2]int{}

    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 1 {
                d[i][j] = 0
                queue = append(queue, [2]int{i, j})
            }
        }
    }

    for len(queue) > 0 {
        x, y := queue[0][0], queue[0][1]
        queue = queue[1:]

        for i := 0; i < 4; i++ {
            u, v := x + dir[i], y + dir[i+1]
            if u >= 0 && u < n && v >= 0 && v < m && d[u][v] == math.MaxInt {
                d[u][v] = d[x][y] + 1
                queue = append(queue, [2]int {u, v})
            }
        }
    }

    bfs := func (val int) bool {    
        visited := make([][]bool, n)
        for i := range visited {
            visited[i] = make([]bool, m)
        }
        queue := [][2]int{[2]int{0, 0}}
        visited[0][0] = true
        if d[0][0] < val {
            return false
        }

        for len(queue) > 0 {
            x, y := queue[0][0], queue[0][1]
            queue = queue[1:]

            if x == n-1 && y == m-1 {
                return true
            }

            for i := 0; i < 4; i++ {
                u, v := x + dir[i], y + dir[i+1] 
                if u >= 0 && u < n && v >= 0 && v < m && !visited[u][v] && d[u][v] >= val {
                    queue = append(queue, [2]int{u, v})
                    visited[u][v] = true
                }
            }
        }

        return false
    }

    low, high := 0, n+m-2
    ans := 0

    for low <= high {
        mid := (low + high) >> 1
        if bfs(mid) { 
            ans = mid
            low = mid+1
        } else {
            high = mid-1
        }
    }

    return ans
}
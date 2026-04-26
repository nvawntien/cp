type State struct {
    x, y int
    px, py int
}

var mvX = []int{1, -1, 0, 0}
var mvY = []int{0, 0, 1, -1}

func containsCycle(grid [][]byte) bool {
    n := len(grid)
    m := len(grid[0])
    
    visited := make([][]bool, n)
    for i := range visited {
        visited[i] = make([]bool, m)
    }

    var bfs func (int, int) bool
    bfs = func(s, t int) bool {
        queue := []State{}
        queue = append(queue, State{s, t, -1, -1})
        visited[s][t] = true

        for len(queue) > 0 {
            x, y, px, py := queue[0].x, queue[0].y, queue[0].px, queue[0].py
            queue = queue[1:]

            for i := 0; i < 4; i++ {
                u, v := x + mvX[i], y + mvY[i]
                if u < 0 || u >= n || v < 0 || v >= m{
                    continue
                } 

                if grid[u][v] != grid[s][t] {
                    continue
                }

                if u == px && v == py {
                    continue
                }

                if visited[u][v] {
                    return true
                }

                queue = append(queue, State{u, v, x, y})
                visited[u][v] = true
            }
        }

        return false
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if !visited[i][j] {
                if bfs(i, j) {
                    return true
                }
            }
        }
    }

    return false
}
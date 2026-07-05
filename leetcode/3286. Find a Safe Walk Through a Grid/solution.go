func findSafeWalk(grid [][]int, health int) bool {
    n := len(grid)
    m := len(grid[0])

    dist := make([][]int, n)
    for i := range dist {
        dist[i] = make([]int, m)
        for j := range dist[i] {
            dist[i][j] = math.MaxInt
        }
    }

    size := n * m * 2 + 100
    deque := make([][2]int, size)
    head := size / 2
    tail := head

    dist[0][0] = grid[0][0]
    deque[tail] = [2]int {0, 0}
    tail++

    dir := [5]int {0, 1, 0, -1, 0}

    for head < tail {
        r, c := deque[head][0], deque[head][1]
        head++

        if r == n-1 && c == m-1 {
            break
        }

        for i := 0; i < 4; i++ {
            nr, nc := r + dir[i], c + dir[i+1]

            if nr >= 0 && nr < n && nc >= 0 && nc < m {
                if dist[r][c] + grid[nr][nc] < dist[nr][nc] {
                    dist[nr][nc] = dist[r][c] + grid[nr][nc]
                    if grid[nr][nc] == 0 {
                        head--
                        deque[head] = [2]int{nr, nc}
                    } else {
                        deque[tail] = [2]int{nr, nc}
                        tail++
                    }
                }
                
            }
        }
    }

    if dist[n-1][m-1] >= health {
        return false
    }

    return true
}
var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func containsCycle(grid [][]byte) bool {
	n := len(grid)
	if n == 0 {
		return false
	}
	m := len(grid[0])
	visited := make([]bool, n*m)

	var dfs func(x, y, px, py int, color byte) bool
	dfs = func(x, y, px, py int, color byte) bool {
		visited[x*m+y] = true

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]

			if nx >= 0 && nx < n && ny >= 0 && ny < m && grid[nx][ny] == color {
				if nx == px && ny == py {
					continue
				}

				if visited[nx*m+ny] {
					return true
				}

				if dfs(nx, ny, x, y, color) {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i*m+j] {
				if dfs(i, j, -1, -1, grid[i][j]) {
					return true
				}
			}
		}
	}

	return false
}
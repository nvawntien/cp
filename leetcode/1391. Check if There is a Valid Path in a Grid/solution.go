type State struct {
	x, y int
}

var dirs = map[int][][]int{
	1: {{0, -1}, {0, 1}},
	2: {{-1, 0}, {1, 0}},
	3: {{0, -1}, {1, 0}},
	4: {{0, 1}, {1, 0}},
	5: {{0, -1}, {-1, 0}},
	6: {{0, 1}, {-1, 0}},
}

func hasValidPath(grid [][]int) bool {
	n := len(grid)
	m := len(grid[0])

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	queue := []State{}
	queue = append(queue, State{0, 0})
	visited[0][0] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		x, y := curr.x, curr.y

		if x == n-1 && y == m-1 {
			return true
		}

		for _, d := range dirs[grid[x][y]] {
			u, v := x+d[0], y+d[1]

			if u >= 0 && u < n && v >= 0 && v < m && !visited[u][v] {
				connect := false
				for _, nd := range dirs[grid[u][v]] {
					if nd[0] == -d[0] && nd[1] == -d[1] {
						connect = true
						break
					}
				}

				if connect {
					queue = append(queue, State{u, v})
					visited[u][v] = true
				}
			}
		}
	}

	return false
}
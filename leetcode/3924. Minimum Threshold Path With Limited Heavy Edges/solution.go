func minimumThreshold(n int, edges [][]int, source int, target int, k int) int {
	if source == target {
		return 0
	}

	adj := make([][][2]int, n)

	for i := range adj {
		adj[i] = make([][2]int, 0)
	}

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	low, high := 0, int(1e9)

	var bfs func(mid int) bool
	bfs = func(mid int) bool {
		front := make([]int, 0)
		back := make([]int, 0)
		back = append(back, source)
		d := make([]int, n)
		for i := range d {
			d[i] = math.MaxInt32
		}
		d[source] = 0

		for len(front) > 0 || len(back) > 0 {
			var u int
			if len(front) > 0 {
				u = front[len(front)-1]
				front = front[:len(front)-1]
			} else {
				u = back[0]
				back = back[1:]
			}

			if d[u] > k {
				continue
			}

			for _, edge := range adj[u] {
				v, w := edge[0], edge[1]
				var cost int
				if w <= mid {
					cost = 0
				} else {
					cost = 1
				}

				if d[u]+cost < d[v] {
					d[v] = d[u] + cost
					if cost == 0 {
						front = append(front, v)
					} else {
						back = append(back, v)
					}
				}
			}
		}
		return d[target] <= k
	}

	threshold := -1

	for low <= high {
		mid := (low + high) >> 1
		if bfs(mid) {
			threshold = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return threshold
}
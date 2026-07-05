func findMaxPathScore(edges [][]int, online []bool, k int64) int {
    n := len(online)

    low, high := 0, 0
    adj := make([][][2]int, n)
    for _, e := range edges {
        u, v, c := e[0], e[1], e[2]
        if !online[u] || !online[v] {
            continue
        }
        high = max(high, c)
        adj[u] = append(adj[u], [2]int {c, v})
    }

    check := func (tar int) bool {
        dp := make([]int64, n)
        for i := 1; i < n; i++ {
            dp[i] = math.MaxInt64
        }

        pq := &PriorityQueue{}
        heap.Init(pq)
        heap.Push(pq, [2]int64{0, 0})

        for pq.Len() > 0 {
            it := heap.Pop(pq).([2]int64)
            d, u := it[0], it[1]

            if d > dp[u] {
                continue
            }

            if u == int64(n-1) {
                return true
            }

            for _, e := range adj[u] {
                c, v := e[0], e[1]
                if c >= tar {
                    nd := dp[u] + int64(c)
                    if nd <= k && nd < dp[v] {
                        dp[v] = nd
                        heap.Push(pq, [2]int64{dp[v], int64(v)})
                    }
                }
            }
        }

        return false
    }

    ans := -1
    for low <= high {
        mid := (low + high) >> 1
        if check(mid) {
            ans = mid
            low = mid+1
        } else {
            high = mid-1
        }
    }


    return ans
}

type PriorityQueue [][2]int64

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.([2]int64)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)

	item := old[n-1]
	*pq = old[:n-1]

	return item
}

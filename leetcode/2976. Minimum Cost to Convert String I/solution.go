func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
    edges := make([][][2]int, 26)

    for i := range cost {
        edges[original[i]-'a'] = append(edges[original[i]-'a'], [2]int{cost[i], int(changed[i]-'a')})        
    }

    dist := make([][]int64, 26)

    for i := 0; i < 26; i++ {
        dist[i] = dijkstra(i, 26, edges)
    }

    ans := int64(0)

    for i := range source {
        a, b := source[i]-'a', target[i]-'a'
        if a == b {
            continue
        }

        if dist[a][b] == math.MaxInt64 {
            return -1
        }

        ans += dist[a][b]
    }

    return ans
}

func dijkstra(start, n int, edges [][][2]int) []int64 {
    dist := make([]int64, n)
    
    for i := 0; i < n; i++ {
        dist[i] = math.MaxInt64
    }

    dist[start] = 0

    Q := NewPQ[[2]int64] (func (a, b [2]int64) bool {
        if a[0] != b[0] {
            return a[0] < b[0]
        }
        return a[1] < b[1]
    })

    Q.Push([2]int64{dist[start], int64(start)})

    for !Q.Empty() {
        curr := Q.Pop()
        d, u := curr[0], int(curr[1])

        if d > dist[u] {
            continue
        }

        for _, e := range edges[u] {
            w, v := e[0], e[1]
            if dist[u] + int64(w) < dist[v] {
                dist[v] = dist[u] + int64(w)
                Q.Push([2]int64{dist[v], int64(v)})
            }
        }
    }

    return dist
}

type PriorityQueue[T any] struct { data []T; less func(a, b T) bool }

func NewPQ[T any](less func(a, b T) bool) *PriorityQueue[T] { return &PriorityQueue[T]{less: less} }

func (pq *PriorityQueue[T]) Push(v T) { pq.data = append(pq.data, v); pq.up(len(pq.data) - 1) }

func (pq *PriorityQueue[T]) Pop() T {
	n := len(pq.data) - 1; v := pq.data[0]
	pq.data[0], pq.data[n] = pq.data[n], pq.data[0]; pq.data = pq.data[:n]; pq.down(0); return v
}

func (pq *PriorityQueue[T]) Top() T { return pq.data[0] }

func (pq *PriorityQueue[T]) Len() int { return len(pq.data) }

func (pq *PriorityQueue[T]) Empty() bool { return len(pq.data) == 0 }

func (pq *PriorityQueue[T]) up(i int) {
	for i > 0 {
		p := (i - 1) >> 1
		if !pq.less(pq.data[i], pq.data[p]) { break }
		pq.data[i], pq.data[p] = pq.data[p], pq.data[i]; i = p
	}
}

func (pq *PriorityQueue[T]) down(i int) {
	n := len(pq.data)
	for {
		left := (i << 1) + 1; if left >= n { break }
		j := left; right := left + 1
		if right < n && pq.less(pq.data[right], pq.data[left]) { j = right }
		if !pq.less(pq.data[j], pq.data[i]) { break }
		pq.data[i], pq.data[j] = pq.data[j], pq.data[i]; i = j
	}
}
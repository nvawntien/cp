func minCost(n int, edges [][]int) int {
    adj := make([][][2]int, n)

    for _, edge := range edges {
        adj[edge[0]] = append(adj[edge[0]], [2]int{edge[2], edge[1]})
        adj[edge[1]] = append(adj[edge[1]], [2]int{edge[2] * 2, edge[0]})        
    }

    return dijkstra(0, n-1, adj)
}

func dijkstra(start, end int, adj [][][2]int) int {
    dist := make([]int, end+1)
    for i := range dist {
        dist[i] = int(1e9)
    }

    dist[start] = 0
    
    Q := NewPQ[[2]int] (func (a, b [2]int) bool {
        if a[0] != b[0] {
            return a[0] < b[0]
        }

        return a[1] < b[1]
    })

    Q.Push([2]int{0, start})

    for !Q.Empty() {
        edge := Q.Pop()
        u := edge[1]
        d := edge[0]

        if d > dist[u] {
            continue
        }

        for _, e := range adj[u] {
            w, v := e[0], e[1]
            if dist[u] + w < dist[v] {
                dist[v] = dist[u] + w
                Q.Push([2]int{dist[v], v})
            }
         }
    }

    if dist[end] == int(1e9) {
        return -1
    }

    return dist[end]
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
/*
.   author: n.vantien
.   from: University of Engineering and Technology - VNU
*/

package main

import (
	"bufio"
	. "fmt"
	"math/bits"
	"os"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
	m   = 0
)

type SegmentTree []struct {
	lazy  int
	color int
}

func (st SegmentTree) Pushdown(id int) {
	if st[id].lazy == 0 {
		return
	}

	st[id<<1].color = 1 << st[id].lazy
	st[id<<1].lazy = st[id].lazy
	st[id<<1|1].color = 1 << st[id].lazy
	st[id<<1|1].lazy = st[id].lazy
	st[id].lazy = 0
}

func (st SegmentTree) Build(id, l, r int, tour []int) {
	if l == r {
		st[id].color = 1 << tour[l]
		return
	}

	mid := (l + r) >> 1
	st.Build(id<<1, l, mid, tour)
	st.Build(id<<1|1, mid+1, r, tour)

	st[id].color = st[id<<1].color | st[id<<1|1].color
}

func (st SegmentTree) Update(id, l, r, u, v, val int) {
	if u > r || v < l {
		return
	}

	if u <= l && r <= v {
		st[id].lazy = val
		st[id].color = 1 << val
		return
	}

	st.Pushdown(id)
	mid := (l + r) >> 1

	st.Update(id<<1, l, mid, u, v, val)
	st.Update(id<<1|1, mid+1, r, u, v, val)

	st[id].color = st[id<<1].color | st[id<<1|1].color
}

func (st SegmentTree) Get(id, l, r, u, v int) int {
	if l > v || u > r {
		return 0
	}

	if u <= l && r <= v {
		return st[id].color
	}

	st.Pushdown(id)
	mid := (l + r) >> 1

	return st.Get(id<<1, l, mid, u, v) | st.Get(id<<1|1, mid+1, r, u, v)
}

func dfs(u, p int, start, end, tour []int, edges [][]int) {
	m++
	tour[m] = u
	start[u] = m

	for _, v := range edges[u] {
		if v == p {
			continue
		}

		dfs(v, u, start, end, tour, edges)
	}

	end[u] = m
}

func solve() {
	var n, q int
	Fscan(in, &n, &q)

	arr := make([]int, n+1)
	end := make([]int, n+1)
	edges := make([][]int, n+1)
	start := make([]int, n+1)
	tour := make([]int, n+1)

	for i := 1; i <= n; i++ {
		Fscan(in, &arr[i])
	}

	for i := 1; i < n; i++ {
		var u, v int
		Fscan(in, &u, &v)
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}

	dfs(1, 0, start, end, tour, edges)

	for i := 1; i <= n; i++ {
		tour[i] = arr[tour[i]]
	}

	st := make(SegmentTree, 4*n)
	st.Build(1, 1, n, tour)

	for i := 1; i <= q; i++ {
		var t, u, x int
		Fscan(in, &t)

		if t == 1 {
			Fscan(in, &u, &x)
			st.Update(1, 1, n, start[u], end[u], x)
		} else {
			Fscan(in, &u)
			Fprintln(out, bits.OnesCount(uint(st.Get(1, 1, n, start[u], end[u]))))
		}
	}
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	T = 1
	for t := 1; t <= T; t++ {
		solve()
	}
}

/*
.   author: n.vantien
.   from: University of Engineering and Technology - VNU
*/

package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

const (
	maxN = 1000006
	oo   = math.MaxInt64
	LOG  = 20
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

type Pair struct{ x, y int }

func dfs(u, p int, adj [][]int, maxd []Pair, depth []int, parent []int) {
	for _, v := range adj[u] {
		if v != p {
			parent[v] = u
			depth[v] = depth[u] + 1
			dfs(v, u, adj, maxd, depth, parent)

			if maxd[v].x+1 > maxd[u].x {
				maxd[u].y = maxd[u].x
				maxd[u].x = maxd[v].x + 1
			} else if maxd[v].x+1 > maxd[u].y {
				maxd[u].y = maxd[v].x + 1
			}
		}
	}

}

func preprocess(n int, depth []int, parent []int, ancestor [][LOG]int, binup [][LOG]int, maxd []Pair) {
	for u := 1; u <= n; u++ {
		binup[u][0] = maxd[parent[u]].x - depth[parent[u]]
		if maxd[u].x+1 == maxd[parent[u]].x {
			binup[u][0] = maxd[parent[u]].y - depth[parent[u]]
		}
		ancestor[u][0] = parent[u]
	}

	for j := 1; j <= LOG-1; j++ {
		for u := 1; u <= n; u++ {
			ancestor[u][j] = ancestor[ancestor[u][j-1]][j-1]
			binup[u][j] = max(binup[u][j-1], binup[ancestor[u][j-1]][j-1])
		}
	}
}

func findFarthest(u, k int, depth []int, ancestor [][LOG]int, binup [][LOG]int, maxd []Pair) int {
	ini := depth[u]
	k = min(k, depth[u])
	distance := maxd[u].x - depth[u]

	for j := LOG - 1; j >= 0; j-- {
		if k >= (1 << j) {
			distance = max(distance, binup[u][j])
			u = ancestor[u][j]
			k -= (1 << j)
		}
	}

	return distance + ini
}

func solve() {
	var n, q int
	Fscan(in, &n)

	adj := make([][]int, n+1)
	depth := make([]int, n+1)
	maxd := make([]Pair, n+1)
	parent := make([]int, n+1)
	binup := make([][LOG]int, n+1)
	ancestor := make([][LOG]int, n+1)

	parent[1] = 1

	for i := 1; i < n; i++ {
		var u, v int
		Fscan(in, &u, &v)

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	dfs(1, 0, adj, maxd, depth, parent)
	preprocess(n, depth, parent, ancestor, binup, maxd)

	Fscan(in, &q)

	for i := 1; i <= q; i++ {
		var u, k int
		Fscan(in, &u, &k)
		Fprint(out, findFarthest(u, k, depth, ancestor, binup, maxd), " ")
	}

	Fprintln(out)
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	Fscan(in, &T)
	for t := 1; t <= T; t++ {
		solve()
	}
}

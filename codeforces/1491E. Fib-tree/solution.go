package main

import (
	"bufio"
	. "fmt"
	"os"
)

const MAXN = 200005

var (
	in        *bufio.Reader
	out       *bufio.Writer
	edges     = make([][]int, MAXN)
	fibo      = make([]int, 30)
	fiboIndex = make([]int, MAXN)
	size      = make([]int, MAXN)
	removed   = make([]bool, MAXN)
)

func dfs(u, p int) int {
	size[u] = 1
	for _, v := range edges[u] {
		if v == p || removed[v] {
			continue
		}
		size[u] += dfs(v, u)
	}
	return size[u]
}

func calc(u, p, k int, cut *[2]int) bool {
	for _, v := range edges[u] {
		if v == p || removed[v] {
			continue
		}
		if calc(v, u, k, cut) {
			return true
		}
		if size[v] == fibo[k-1] {
			cut[0], cut[1] = v, u
			return true
		}
		if size[v] == fibo[k-2] {
			cut[0], cut[1] = u, v
			return true
		}
	}
	return false
}

func check(u, n int) bool {
	if n <= 3 {
		return true
	}
	k := fiboIndex[n]
	if k == -1 {
		return false
	}
	dfs(u, -1)
	var cut [2]int
	if !calc(u, -1, k, &cut) {
		return false
	}
	a, b := cut[0], cut[1]

	removed[b] = true
	left := check(a, fibo[k-1])
	removed[b] = false

	removed[a] = true
	right := check(b, fibo[k-2])
	removed[a] = false

	return left && right
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int

	Fscan(in, &n)

	for i := 0; i <= n; i++ {
		edges[i] = make([]int, 0)
	}

	for i := 1; i < n; i++ {
		var u, v int
		Fscan(in, &u, &v)
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}

	fibo[0], fibo[1] = 1, 1
	fiboIndex[1] = 1
	for i := 2; i < 30; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
		if fibo[i] < MAXN {
			fiboIndex[fibo[i]] = i
		}
	}

	if fiboIndex[n] == 0 {
		Fprintln(out, "NO")
		return
	}

	if check(1, n) {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}
}

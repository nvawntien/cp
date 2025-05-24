/*
.   author: n.vantien
.   from: University of Engineering and Technology - VNU
*/

package main

import (
	"bufio"
	. "fmt"
	"os"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

func dfs(u int, ans *int, road [][]int, citizens []int) (int, int) {
	leaf := 0
	if len(road[u]) == 0 {
		leaf++
	}
	citizen := citizens[u]
	for _, v := range road[u] {
		l, c := dfs(v, ans, road, citizens)
		leaf += l
		citizen += c
	}
	*ans = max(*ans, (citizen-1)/leaf+1)
	return leaf, citizen
}

func solve() {
	var n int
	Fscan(in, &n)

	road := make([][]int, n+1)
	citizens := make([]int, n+1)
	sum := 0

	for i := 2; i <= n; i++ {
		var p int
		Fscan(in, &p)
		road[p] = append(road[p], i)
	}

	for i := 1; i <= n; i++ {
		Fscan(in, &citizens[i])
		sum += citizens[i]
	}

	if sum == 0 {
		Fprintln(out, 0)
		return
	}

	ans := 0
	dfs(1, &ans, road, citizens)

	Fprint(out, ans)
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

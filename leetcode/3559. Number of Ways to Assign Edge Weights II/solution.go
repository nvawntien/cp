func assignEdgeWeights(edges [][]int, queries [][]int) []int {
    n := len(edges)+2
    graph := make([][]int, n)

    for _, e := range edges {
        u, v := e[0], e[1]
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }

    log := bits.Len(uint(n))
    up  := make([][]int, n)
    depth := make([]int, n)
    mod  := int(1e9 + 7)

    for i := range up {
        up[i] = make([]int, log)
    }

    var dfs func (int) 
    dfs = func (u int) {
        for _, v := range graph[u] {
            if v == up[u][0] {
                continue
            }

            depth[v] = depth[u] + 1
            up[v][0] = u

            for j := 1; j < log; j++ {
                up[v][j] = up[up[v][j-1]][j-1]
            }

            dfs(v)
        }
    }

    lca := func (u, v int) int {
        if depth[u] != depth[v] {
            if depth[u] < depth[v] {
                u, v = v, u
            } 
            k := depth[u] - depth[v]
            for j := 0; (1<<j) <= k; j++ {
                if k>>j&1 == 1 {
                    u = up[u][j]
                }
            }
        }

        if u == v {
            return u
        }

        k := bits.Len(uint(depth[u]))-1
        for j := k; j >= 0; j-- {
            if up[u][j] != up[v][j] {
                u = up[u][j]
                v = up[v][j]
            }
        }

        return up[u][0]
    }

    dist := func (u, v int) int {
        p := lca(u, v)
        return depth[u] + depth[v] - 2 * depth[p]
    }

    pow := func (a, b int) int {
        ans := 1
        for b > 0 {
            if b&1 == 1 {
                ans = ans * a % mod
            }
            a = a * a % mod
            b >>= 1
        }
        return ans
    }

    dfs(1)

    m := len(queries)
    ans := make([]int, m)

    for i, q := range queries {
        u, v := q[0], q[1]
        if u == v {
            ans[i] = 0
            continue
        }

        ans[i] = pow(2, dist(u, v)-1)
    }

    return ans
}
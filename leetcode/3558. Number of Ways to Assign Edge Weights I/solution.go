func assignEdgeWeights(edges [][]int) int {
    n := len(edges)+1
    adj := make([][]int, n+1)

    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }

    var dfs func(int, int) int
    dfs = func (p, u int) int {
        depth := 0
        for _, v := range adj[u] {
            if v == p {
                continue
            }

            depth = max(depth, 1 + dfs(u, v))
        }
        return depth
    } 

    b := dfs(-1, 1)-1
    a := 2
    mod := int(1e9 + 7)
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
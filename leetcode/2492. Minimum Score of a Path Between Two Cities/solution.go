func minScore(n int, roads [][]int) int {
    graph := make([][][2]int, n+1)
    for _, e := range roads {
        u, v, d := e[0], e[1], e[2]
        graph[u] = append(graph[u], [2]int {v, d})
        graph[v] = append(graph[v], [2]int {u, d})
    }

    visited := make([]bool, n+1)
    var dfs func (p, u int) int    
    dfs = func (p, u int) int {
        var ans int = math.MaxInt
        visited[u] = true
        for _, e := range graph[u] {
            v, d := e[0], e[1]
            ans = min(ans, d)
            if v != p && !visited[v] {
                ans = min(ans, dfs(u, v))
            }
        }
        return ans
    }

    return dfs(-1, 1)
}
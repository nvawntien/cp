func countCompleteComponents(n int, edges [][]int) int {
    graph := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }

    visited := make([]bool, n)
    var dfs func (u int) (int, int)

    dfs = func (u int) (int, int) {
        node, edge := 1, 0
        edge += len(graph[u])
        visited[u] = true
        for _, v := range graph[u] {
            if visited[v] {
                continue
            }

            n, e := dfs(v)
            node += n
            edge += e
        }
        return node, edge
    }

    ans := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            node, edge := dfs(i)
            if node * (node-1) == edge {
                ans++
            }
        } 
    }

    return ans
}
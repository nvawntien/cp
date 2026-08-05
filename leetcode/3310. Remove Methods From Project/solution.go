func remainingMethods(n int, k int, invocations [][]int) []int {
    graph := make([][]int, n)

    for _, inv := range invocations {
        u, v := inv[0], inv[1]
        graph[u] = append(graph[u], v)
    }

    bug := make([]bool, n)

    var find func(int)
    find = func(u int) {
        bug[u] = true
        for _, v := range graph[u] {
            if !bug[v] {
                find(v)
            }
        }
    }

    find(k)

    ans := make([]int, 0, n)
    check := false
    for _, inv := range invocations {
        u, v := inv[0], inv[1]
        if !bug[u] && bug[v] {
            check = true
            break
        }
    }

    if check {
        for i := 0; i < n; i++ {
            ans = append(ans, i)
        }
        return ans
    }

    for i := 0; i < n; i++ {
        if bug[i] {
            continue
        }

        ans = append(ans, i)
    }

    return ans   
}
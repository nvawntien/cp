func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
    const MOD = 1e9 + 7
    hFences = append(hFences, 1)
    hFences = append(hFences, m)
    vFences = append(vFences, 1)
    vFences = append(vFences, n)

    sort.Slice(hFences, func (i, j int) bool {
        return hFences[i] < hFences[j]
    })

    sort.Slice(vFences, func (i, j int) bool {
        return vFences[i] < vFences[j]
    })

    N, M := len(hFences), len(vFences)

    mp := make(map[int]bool)

    for i := 0; i < N; i++ {
        for j := 0; j < i; j++ {
            mp[hFences[i] - hFences[j]] = true
        }
    }

    edgeMax := -1

    for i := 0; i < M; i++ {
        for j := 0; j < i; j++ {
            if mp[vFences[i] - vFences[j]] {
                edgeMax = max(edgeMax, vFences[i] - vFences[j])
            }
        }
    }

    if edgeMax == -1 {
        return -1
    }

    return (edgeMax % MOD)  * (edgeMax % MOD) % MOD
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
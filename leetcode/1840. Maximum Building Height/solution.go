func maxBuilding(n int, res [][]int) int {
    res = append(res, []int{1, 0})
    sort.Slice(res, func (i, j int) bool {
        return res[i][0] < res[j][0]
    })

    m := len(res)

    for i := 1; i < m; i++ {
        res[i][1] = min(res[i][1], res[i-1][1] + res[i][0] - res[i-1][0])
    }

    for i := m-2; i >= 0; i-- {
        res[i][1] = min(res[i][1], res[i+1][1] + res[i+1][0] - res[i][0])
    }

    ans := 0

    for i := 0; i < m-1; i++ {
        d := res[i+1][0] - res[i][0]
        h1 := res[i][1]
        h2 := res[i+1][1]
        ans = max(ans, (h1+h2+d) >> 1)
    }

    ans = max(ans, res[m-1][1] + n-res[m-1][0])
    return ans
}
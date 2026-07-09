func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }

    for i := 1; i < n; i++ {
        if nums[i] - nums[i-1] <= maxDiff {
            parent[i] = parent[i-1]
        }
    }

    ans := make([]bool, len(queries))

    for i, q := range queries {
        if parent[q[0]] != parent[q[1]] {
            ans[i] = false
        } else {
            ans[i] = true
        }
    }

    return ans
}
func minAbsDiff(grid [][]int, k int) [][]int {
    n := len(grid)
    m := len(grid[0])

    ans := make([][]int, n-k+1)
    for i := range ans {
        ans[i] = make([]int, m-k+1)
    }

    for i := 0; i <= n-k; i++ {
        for j := 0; j <= m-k; j++ {
            arr := []int{}
            for u := i; u < i+k; u++ {
                for v := j; v < j+k; v++ {
                    arr = append(arr, grid[u][v])
                }
            }

            sort.Ints(arr)
            diff := int(1e9)

            for t := 1; t < len(arr); t++ {
                if arr[t] != arr[t-1] {
                    if arr[t] - arr[t-1] < diff {
                        diff = arr[t] - arr[t-1]
                    }
                }
            }

            if diff != int(1e9) {
                ans[i][j] = diff
            } else {
                ans[i][j] = 0
            }
        }
    }

    return ans
}
func maxJumps(arr []int, d int) int {
    n := len(arr)
    
    var dfs func (i int) int
    mem := make(map[int]int)

    dfs = func (i int) int {
        if val, ok := mem[i]; ok {
            return val
        }

        res := 1
        for j := i+1; j < min(i+d+1, n); j++ {
            if arr[j] >= arr[i] {
                break
            }
            res = max(res, 1+dfs(j))
        }

        for j := i-1; j >= max(i-d, 0); j-- {
            if arr[j] >= arr[i] {
                break
            }
            res = max(res, 1+dfs(j))
        }

        mem[i] = res
        return res
    }

    ans := 0

    for i := 0; i < n; i++ {
        ans = max(ans, dfs(i))
    }

    return ans
} 
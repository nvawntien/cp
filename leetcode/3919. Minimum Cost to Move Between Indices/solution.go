func minCost(nums []int, queries [][]int) []int {
    n := len(nums)
    prel := make([]int, n)
    prer := make([]int, n)

    for i := 1; i < n; i++ {
        var cost int
        if i == 1 {
            cost = 1
        } else {
            dl := nums[i-1] - nums[i-2]
            dr := nums[i] - nums[i-1]
            if dr < dl {
                cost = 1
            } else {
                cost = dr
            }
        }

        prel[i] = prel[i-1] + cost
    }
    
    for i := n-2; i >= 0; i-- {
        var cost int
        if i == n-2 {
            cost = 1
        } else {
            dr := nums[i+2] - nums[i+1]
            dl := nums[i+1] - nums[i]
            if dl <= dr {
                cost = 1
            } else {
                cost = dl
            }
        }

        prer[i] = prer[i+1] + cost
    }

    ans := make([]int, len(queries))
    
    for i, q := range queries {
        u, v := q[0], q[1]
        if u < v {
            ans[i] = prel[v] - prel[u]        
        } else {
            ans[i] = prer[v] - prer[u]
        }
    }

    return ans
}
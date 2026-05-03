type Pair struct {
    X int
    D int
}

func maxFixedPoints(nums []int) int {
    n := len(nums)
    valid := make([]Pair, 0)

    for i := 0; i < n; i++ {
        if nums[i] <= i {
            valid = append(valid, Pair{X: nums[i], D: i - nums[i]})
        }
    }

    sort.Slice(valid, func (i, j int) bool {
        if valid[i].X != valid[j].X {
            return valid[i].X < valid[j].X
        }
        return valid[i].D > valid[j].D
    })

    dp := make([]int, 0)

    for _, p := range valid {
        d := p.D
        idx := sort.Search(len(dp), func (i int) bool {
            return dp[i] > d
        })

        if idx == len(dp) {
            dp = append(dp, d)
        } else {
            dp[idx] = d
        }
    }

    return len(dp)
}
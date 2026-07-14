func subsequencePairCount(nums []int) int {
    const MOD = 1_000_000_007
    const MAX = 201

    gcd := func (a, b int) int {
        for b != 0 {
            a, b = b, a % b
        }
        return a
    }

    dp := make([][]int, MAX)
    for i := range dp {
        dp[i] = make([]int, MAX)
    }
    dp[0][0] = 1

    for _, x := range nums {
        next := make([][]int, MAX)
        for i := range next {
            next[i] = make([]int, MAX)
        }

        for g1 := 0; g1 < MAX; g1++ {
            for g2 := 0; g2 < MAX; g2++ {
                if g := dp[g1][g2]; g > 0 {
                    next[g1][g2] = (next[g1][g2] + g) % MOD
                    
                    ng1 := x
                    if g1 != 0 {
                        ng1 = gcd(g1, x)
                    }
                    next[ng1][g2] = (next[ng1][g2] + g) % MOD
                    
                    ng2 := x
                    if g2 != 0 {
                        ng2 = gcd(g2, x)
                    }
                    next[g1][ng2] = (next[g1][ng2] + g) % MOD
                }
            }
        }
        dp = next
    }

    ans := 0
    for g := 1; g < MAX; g++ {
        ans = (ans + dp[g][g]) % MOD
    }

    return ans
}
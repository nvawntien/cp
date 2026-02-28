func concatenatedBinary(n int) int {
    bitLen := 0
    ans := 0
    const MOD = 1e9 + 7

    for i := 1; i <= n; i++ {
        if i & (i-1) == 0 {
            bitLen++
        }
        ans = (ans << bitLen | i) % MOD
    }

    return ans
}
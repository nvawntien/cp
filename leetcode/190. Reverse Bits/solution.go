func reverseBits(n int) int {
    ans := 0
    two := 1 << 31
    for n > 0 {
        ans = ans + (n & 1) * two
        two >>= 1
        n >>= 1
    }
    return ans
}
func hasAlternatingBits(n int) bool {
    bit := n & 1
    n >>= 1
    
    for n > 0 {
        ans := bit ^ n & 1
        if ans != 1 {
            return false
        }
        bit = n & 1
        n >>= 1
    }

    return true
}
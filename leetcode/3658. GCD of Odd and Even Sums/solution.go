func gcdOfOddEvenSums(n int) int {
    odd := (((n-1)<<1 + 2) * n) >> 1
    even := (((n-1)<<1 + 4) * n) >> 1

    for even != 0 {
        odd, even = even, odd % even
    }

    return odd
}
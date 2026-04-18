func mirrorDistance(n int) int {
    rev := 0
    cur := n
    for cur > 0 {
        rev = rev * 10 + cur % 10
        cur = cur / 10
    }

    ans := rev - n
    mask := ans >> 31
    return (ans ^ mask) - mask
}
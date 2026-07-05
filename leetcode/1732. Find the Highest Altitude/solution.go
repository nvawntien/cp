func largestAltitude(gain []int) int {
    n := len(gain)
    s := 0
    ans := s

    for i := 0; i < n; i++ {
        ans = max(ans, s + gain[i])
        s = s + gain[i]
    }

    return ans
}
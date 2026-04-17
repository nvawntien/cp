func countBinarySubstrings(s string) int {
    cur, pre, ans := 1, 0, 0
    // 0001110
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            cur++
        } else {
            ans += min(cur, pre)
            pre, cur = cur, 1
        }
    }
    return ans + min(cur, pre)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
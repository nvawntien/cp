func hasAllCodes(s string, k int) bool {
    n    := len(s)
    need := 1 << k
    mask := need-1
    seen := make([]bool, need)
    hash := 0

    for i := 0; i < n; i++ {
        hash = ((hash << 1) & mask) | int(s[i] - '0')

        if i >= k-1 && !seen[hash] {
            seen[hash] = true
            need--
            if need == 0 {
                return true
            }
        }
    }

    return false
}
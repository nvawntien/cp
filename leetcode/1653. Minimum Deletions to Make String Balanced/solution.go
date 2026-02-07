func minimumDeletions(s string) int {
    ans := 0
    count := 0

    for i := range s {
        if s[i] == 'b' {
            count++
        } else {
            if ans + 1 < count {
                ans++
            } else {
                ans = count
            }
        }
    }

    return ans
}
func maxActiveSectionsAfterTrade(s string) int {
    prev := 0
    curr := 0
    ans := 0
    cnt := 0
    n := len(s)

    for i := 0; i < n; i++ {
        if s[i] == '1' {
            ans++
            if curr > 0 {
                if prev > 0 {
                    cnt = max(cnt, prev + curr)
                }
                prev = curr
                curr = 0
            }
        } else {
            curr++
        }
    }

    if prev > 0 && curr >  0 {
        cnt = max(cnt, prev + curr)
    }

    return ans + cnt
}
func canReach(s string, minJump int, maxJump int) bool {
    n := len(s)
    if s[n-1] != '0' {
        return false
    }

    reach := 0
    dp := make([]bool, n)
    dp[0] = true

    for i := 0; i < n; i++ {
        if i >= minJump && dp[i-minJump] {
            reach++
        }

        if i > maxJump && dp[i-maxJump-1] {
            reach--
        }

        if s[i] == '0' && reach > 0 {
            dp[i] = true
        }
    }

    return dp[n-1]
}
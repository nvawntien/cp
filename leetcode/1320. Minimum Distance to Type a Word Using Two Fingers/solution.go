const INF = 1_000_000_000

func minimumDistance(word string) int {
    n := len(word)
    dp := make([][]int, n)

    for i := range dp {
        dp[i] = make([]int, 26)
        for j := range dp[i] {
            dp[i][j] = INF
        }
    }

    for j := 0; j < 26; j++ {
        dp[0][j] = 0
    }

    abs := func(x int) int {
        mask := x >> 31
        return (x ^ mask) - mask   
    }

    dist := func (a, b int) int {
        x1, y1 := a / 6, a % 6
        x2, y2 := b / 6, b % 6
        return abs(x1 - x2) + abs(y1 - y2)
    }

    for i := 1; i < n; i++ {
        p := int(word[i-1] - 'A')
        c := int(word[i] - 'A')

        for j := 0; j < 26; j++ {
            if dp[i-1][j] == INF {
                continue
            }

            dp[i][j] = min(dp[i][j], dp[i-1][j] + dist(p, c))
            dp[i][p] = min(dp[i][p], dp[i-1][j] + dist(j, c))
        }
    }

    var ans = INF

    for j := range dp[n-1] {
        ans = min(ans, dp[n-1][j])
    }

    return ans
}
const mod = 1_000_000_007

func pathsWithMaxScore(board []string) []int {
    n := len(board)
    m := len(board[0])

    dp := make([][][2]int, n)
    for i := range dp {
        dp[i] = make([][2]int, m)
    }

    dp[0][0][0] = 1

    for i := 1; i < n; i++ {
        if board[i][0] == 'X' {
            continue
        }
        if dp[i-1][0][0] != 0 {
            dp[i][0][1] = dp[i-1][0][1] + int(board[i][0] - '0')
        }
        dp[i][0][0] = dp[i-1][0][0]

    }   

    for j := 1; j < m; j++ {
        if board[0][j] == 'X' {
            continue
        }
        if dp[0][j-1][0] != 0 {
            dp[0][j][1] = dp[0][j-1][1] + int(board[0][j] - '0')
        }
        dp[0][j][0] = dp[0][j-1][0]
    }

    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if board[i][j] == 'X' {
                continue
            }
            if i == n-1 && j == m-1 {
                dp[i][j][1] = max(dp[i-1][j][1], max(dp[i][j-1][1], dp[i-1][j-1][1]))
                if dp[i-1][j][1] == dp[i][j][1] {
                    dp[i][j][0] = (dp[i][j][0] + dp[i-1][j][0]) % mod
                }
                if dp[i][j-1][1] == dp[i][j][1] {
                    dp[i][j][0] = (dp[i][j][0] + dp[i][j-1][0]) % mod
                }
                if dp[i-1][j-1][1] == dp[i][j][1] {
                    dp[i][j][0] = (dp[i][j][0] + dp[i-1][j-1][0]) % mod
                }
                break
            }

            if dp[i-1][j][0] != 0 {
                dp[i][j][1] = max(dp[i][j][1], dp[i-1][j][1] + int(board[i][j] - '0')) % mod
            }
            if dp[i][j-1][0] != 0 {
                dp[i][j][1] = max(dp[i][j][1], dp[i][j-1][1] + int(board[i][j] - '0')) % mod
            }
            if dp[i-1][j-1][0] != 0 {
                dp[i][j][1] = max(dp[i][j][1], dp[i-1][j-1][1] + int(board[i][j] - '0')) % mod
            }

            if (dp[i-1][j][1] + int(board[i][j] - '0')) % mod == dp[i][j][1] {
                dp[i][j][0] = (dp[i][j][0] + dp[i-1][j][0]) % mod
            }
            if (dp[i][j-1][1] + int(board[i][j] - '0')) % mod == dp[i][j][1] {
                dp[i][j][0] = (dp[i][j][0] + dp[i][j-1][0]) % mod
            }
            if (dp[i-1][j-1][1] + int(board[i][j] - '0')) % mod == dp[i][j][1] {
                dp[i][j][0] = (dp[i][j][0] + dp[i-1][j-1][0]) % mod
            }
        }
    }

    if dp[n-1][m-1][0] == 0 {
        return []int {0, 0}
    }

    return []int {dp[n-1][m-1][1], dp[n-1][m-1][0]}
}

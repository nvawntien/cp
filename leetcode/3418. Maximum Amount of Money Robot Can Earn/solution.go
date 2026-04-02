const MinInf = -1_000_000_000

func maximumAmount(coins [][]int) int {
    n := len(coins)
    m := len(coins[0])

    prev := make([][3]int, m)
    curr := make([][3]int, m)

    for j := 0; j < m; j++ {
        prev[j] = [3]int{MinInf, MinInf, MinInf}
        curr[j] = [3]int{MinInf, MinInf, MinInf}
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            curr[j] = [3]int{MinInf, MinInf, MinInf}
            
            if i == 0 && j == 0 {
                curr[j][2] = coins[i][j]    
                if coins[i][j] < 0 {
                    curr[j][1] = 0
                }
                continue
            }

            for k := 0; k <= 2; k++ {
                if j > 0 {
                    curr[j][k] = max(curr[j][k], curr[j-1][k] + coins[i][j])
                }

                if i > 0 {
                    curr[j][k] = max(curr[j][k], prev[j][k] + coins[i][j])
                }

                if coins[i][j] < 0 && k < 2 {
                     if j > 0 {
                        curr[j][k] = max(curr[j][k], curr[j-1][k+1])
                    }

                    if i > 0 {
                        curr[j][k] = max(curr[j][k], prev[j][k+1])
                    }
                }
            }
        }

        copy(prev, curr)
    }
    
    ans := MinInf

    for k := 0; k <= 2; k++ {
        ans = max(ans, prev[m-1][k])
    }

    return ans
}
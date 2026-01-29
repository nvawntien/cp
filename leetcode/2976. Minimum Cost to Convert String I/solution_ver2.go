func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
   
    var dist [26][26]int
    const INF = math.MaxInt64 

    for i := 0; i < 26; i++ {
        for j := 0; j < 26; j++ {
            if i == j {
                dist[i][j] = 0
            } else {
                dist[i][j] = INF
            }
        }
    }

    for i := range original {
        u := original[i] - 'a'
        v := changed[i] - 'a'
        c := cost[i]
    
        if c < dist[u][v] {
            dist[u][v] = c
        }
    }

    for k := 0; k < 26; k++ {
        for i := 0; i < 26; i++ {
            for j := 0; j < 26; j++ {
                if dist[i][k] != INF && dist[k][j] != INF {
                    if dist[i][k] + dist[k][j] < dist[i][j] {
                        dist[i][j] = dist[i][k] + dist[k][j]
                    }
                }
            }
        }
    }

    var ans int64 = 0
    for i := 0; i < len(source); i++ {
        srcChar := source[i] - 'a'
        tgtChar := target[i] - 'a'

        if srcChar == tgtChar {
            continue
        }

        if dist[srcChar][tgtChar] == INF {
            return -1 
        }

        ans += int64(dist[srcChar][tgtChar])
    }

    return ans
}
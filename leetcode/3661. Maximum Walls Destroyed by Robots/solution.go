type Robot struct {
    pos int
    dist int
}

func maxWalls(robots []int, distance []int, walls []int) int {
    n := len(robots)
    m := len(walls)
    arr := make([]Robot, n)

    for i := range robots {
        arr[i] = Robot{pos: robots[i], dist: distance[i]}
    }

    sort.Slice(arr, func (i, j int) bool {
        return arr[i].pos < arr[j].pos
    })

    sort.Ints(walls)

    cleanWalls := []int{}
    overlapWalls := 0
    i := 0
    j := 0

    for i < m && j < n {
        if walls[i] == arr[j].pos {
            overlapWalls++
            i++
        } else if walls[i] < arr[j].pos {
            cleanWalls = append(cleanWalls, walls[i])
            i++
        } else {
            j++
        }
    }

    for i < m {
        cleanWalls = append(cleanWalls, walls[i])
        i++
    }

    m -= overlapWalls

    countWall := func (L, R int) int {
        if L > R {
            return 0
        }

        low := sort.SearchInts(cleanWalls, L)
        high := sort.SearchInts(cleanWalls, R+1)

        return high - low
    }

    dp := make([][2]int, n)
    dp[0][0] = countWall(arr[0].pos-arr[0].dist, arr[0].pos-1)
    dp[0][1] = 0

    for k := 1; k < n; k++ {
        L0 := arr[k-1].pos + 1
        R0 := L0 + arr[k-1].dist - 1
        L1 := arr[k].pos - arr[k].dist
        R1 := arr[k].pos - 1

        countA := countWall(L0, min(R0, R1))
        countB := countWall(max(L0, L1), R1)
        overlap := countWall(max(L0, L1), min(R0, R1))

        dp[k][0] = max(dp[k-1][0] + countB, dp[k-1][1] + countA + countB - overlap)

        dp[k][1] = max(dp[k-1][0], dp[k-1][1] + countA)               
    }

    end := countWall(arr[n-1].pos+1, arr[n-1].pos + arr[n-1].dist)

    return overlapWalls + max(dp[n-1][0], dp[n-1][1] + end)
}
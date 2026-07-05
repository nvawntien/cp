type Ride struct {
    s int
    d int
}

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    n := len(landStartTime)
    m := len(waterStartTime)

    solve := func(first, second []Ride) int {
        sort.Slice(second, func (i, j int) bool {
            return second[i].s < second[j].s
        })

        sz := len(second)

        pre := make([]int, sz)
        pre[0] = second[0].d
        for i := 1; i < sz; i++ {
            pre[i] = min(pre[i-1], second[i].d)
        }

        suf := make([]int, sz)
        suf[sz-1] = second[sz-1].s + second[sz-1].d
        for i := sz-2; i >= 0; i-- {
            suf[i] = min(suf[i+1], second[i].s + second[i].d)
        }

        ans := math.MaxInt

        for _, r := range first {
            t := r.s + r.d
            idx := sort.Search(sz, func (i int) bool {
                return second[i].s > t
            })

            if idx > 0 {
                ans = min(ans, t + pre[idx-1])
            }

            if idx < sz {
                ans = min(ans, suf[idx])
            }   
        }

        return ans
    }

    land := make([]Ride, n)
    for i := 0; i < n; i++ {
        land[i] = Ride{s: landStartTime[i], d: landDuration[i]}
    }
    
    water := make([]Ride, m)
    for i := 0; i < m; i++ {
        water[i] = Ride{s: waterStartTime[i], d: waterDuration[i]}
    }

    ans := min(solve(land, water), solve(water, land))

    return ans
}
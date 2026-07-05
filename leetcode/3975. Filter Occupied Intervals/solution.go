func filterOccupiedIntervals(occ [][]int, start int, end int) [][]int {
    sort.Slice(occ, func (i, j int) bool {
        return occ[i][0] < occ[j][0]
    })

    var merge [][]int
    curr := occ[0]
    
    for i := range occ {
        next := occ[i]
        if next[0] <= curr[1]+1 {
            curr[1] = max(curr[1], next[1])
        } else {
            merge = append(merge, curr)
            curr = next
        }
    }

    merge = append(merge, curr)

    var res [][]int
    for _, m := range merge {
        s, e := m[0], m[1]
        if s > end || e < start {
            res = append(res, []int {s, e})
            continue
        }

        if s < start && end < e {
            res = append(res, []int {s, start-1})
            res = append(res, []int {end+1, e})
            continue
        }

        if s < start && end >= e {
            res = append(res, []int {s, start-1})
            continue
        }

        if start <= s && end < e {
            res = append(res, []int {end+1, e})
            continue
        }
    }

    return res
}
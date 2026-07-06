func removeCoveredIntervals(intervals [][]int) int {
    n := len(intervals)
    sort.Slice(intervals, func (i, j int) bool {
        if intervals[i][0] != intervals[j][0] {
            return intervals[i][0] < intervals[j][0]
        }
        return intervals[i][1] > intervals[j][1]
    })

    l, r := -1, -1
    cnt := 0
    for i := 0; i < n; i++ {
        if intervals[i][0] >= l && intervals[i][1] <= r {
            cnt++
        }

        l = max(intervals[i][0], l)
        r = max(intervals[i][1], r)
    }

    return n - cnt
}
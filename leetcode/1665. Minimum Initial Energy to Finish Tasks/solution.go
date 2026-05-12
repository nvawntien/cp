func minimumEffort(tasks [][]int) int {
    sort.Slice(tasks, func (i, j int) bool {
        return (tasks[i][1] - tasks[i][0]) >= (tasks[j][1] - tasks[j][0])
    })

    var check func(int) bool
    check = func (mid int) bool {
        for _, task := range tasks {
            if task[1] > mid {
                return false
            }
            mid -= task[0]
        }
        return mid >= 0
    }

    low, high := 0, int(1e9)
    ans := low

    for low <= high {
        mid := (low + high) >> 1
        if check(mid) {
            ans = mid
            high = mid-1
        } else {
            low = mid+1
        }
    }

    return ans
}
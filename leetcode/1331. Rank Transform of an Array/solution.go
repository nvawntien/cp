type Arr struct {
    val int
    idx int
}

func arrayRankTransform(arr []int) []int {
    n := len(arr)
    if n == 0 {
        return []int{}
    }
    nums := make([]Arr, n)
    for i, v := range arr {
        nums[i] = Arr{val: v, idx: i}
    }    

    sort.Slice(nums, func (i, j int) bool {
        return nums[i].val < nums[j].val
    })

    pos := make([]int, n)
    idx := 1

    pos[nums[0].idx] = idx

    for i := 1; i < n; i++ {
        if nums[i].val != nums[i-1].val {
            idx++
        }
        pos[nums[i].idx] = idx
    }

    return pos
}
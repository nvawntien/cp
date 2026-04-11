func minimumDistance(nums []int) int {
    idx := make(map[int][]int)

    for i := range nums {
        idx[nums[i]] = append(idx[nums[i]], i)
    }

    res := 1_000_000_000

    for _, indices := range idx {
        if len(indices) < 3 {
            continue
        }

        for i := 2; i < len(indices); i++ {
            res = min(res, indices[i] - indices[i-2])
        }
    }

    if res == 1_000_000_000 {
        return -1
    }

    return res << 1
}
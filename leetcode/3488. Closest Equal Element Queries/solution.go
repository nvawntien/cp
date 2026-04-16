func solveQueries(nums []int, queries []int) []int {
    indices := make(map[int][]int)
    m := len(nums)

    for i := range nums {
        indices[nums[i]] = append(indices[nums[i]], i)
    }

    ans := make([]int, len(queries))

    for i, q := range queries {
        idx := indices[nums[q]]
        n := len(idx)
        
        if n == 1 {
            ans[i] = -1
            continue
        }

        pos, _ := slices.BinarySearch(idx, q)
        mid := idx[pos]
        left := (mid - idx[(pos-1+n) % n] + m) % m
        right := (idx[(pos+1) % n] - mid + m) % m

        ans[i] = min(left, right)
    }

    return ans
}
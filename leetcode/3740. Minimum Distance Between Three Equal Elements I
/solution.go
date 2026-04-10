func minimumDistance(nums []int) int {
    // |i - j| + |j - k| + |k - i|
    // i < j < k
    // j - i + k - j + k - i
    // 2k - 2i

    n := len(nums)
    arr := make([][]int, n+1)

    for i := range nums {
        arr[nums[i]] = append(arr[nums[i]], i)
    }

    res := 1_000_000

    for _, indices := range arr {
        if len(indices) < 3 {
            continue
        }

        for i := 2; i < len(indices); i++ {
            res = min(res, indices[i] - indices[i-2])
        }
    }

    if res == 1_000_000 {
        return -1
    }

    return res << 1
}
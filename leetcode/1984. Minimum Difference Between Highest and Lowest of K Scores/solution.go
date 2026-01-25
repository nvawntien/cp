func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)
    ans := int(1e9)
    for i := k-1; i < len(nums); i++ {
        ans = min(ans, nums[i] - nums[i-k+1])
    }

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
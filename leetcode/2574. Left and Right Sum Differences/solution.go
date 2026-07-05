func leftRightDifference(nums []int) []int {
    n := len(nums)
    pre := make([]int, n)
    sum := nums[0]

    for i := 1; i < n; i++ {
        sum += nums[i]
        pre[i] = pre[i-1] + nums[i-1]
    }

    ans := make([]int, n)

    abs := func (x int) int {
        mask := x >> 31
        return (x ^ mask) - mask
    }

    for i := 0; i < n; i++ {
        ans[i] = abs(pre[i] - (sum - pre[i] - nums[i]))
    }

    return ans
}
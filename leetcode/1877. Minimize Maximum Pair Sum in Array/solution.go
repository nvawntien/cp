func minPairSum(nums []int) int {
    sort.Slice(nums, func (i, j int) bool {
        return nums[i] < nums[j]
    })

    n := len(nums)    
    ans := 0
    
    for i := 0; i < n/2; i++ {
        ans = max(ans, nums[i] + nums[n-i-1])    
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
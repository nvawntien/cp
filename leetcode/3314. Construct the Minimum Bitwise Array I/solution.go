func minBitwiseArray(nums []int) []int {
    ans := make([]int, len(nums))

    for i := range nums {
        ans[i] = -1
        for k := 0; k <= nums[i]; k++ {
            if (k | (k+1)) == nums[i] {
                ans[i] = k
                break
            }
        } 
    }

    return ans
}
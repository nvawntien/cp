func maxSum(nums []int, k int, mul int) int64 {
    sort.Ints(nums)
    n := len(nums)

    sum := int64(0)
    
    for i := n-1; i >= n-k; i-- {
        if mul > 0 {
            sum += int64(nums[i] * mul)
            mul--
        } else {
            sum += int64(nums[i])
        }
    }

    return sum
}
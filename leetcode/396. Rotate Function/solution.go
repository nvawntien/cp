func maxRotateFunction(nums []int) int {
    n := len(nums)
    // F(k) =  F(k-1) + S - n * nums[n-k]
    
    sum := 0
    F := 0

    for i, num := range nums {
        sum += num
        F += i * num
    }

    ans := F

    for k := 1; k < n; k++ {
        curr := F + sum - n * nums[n-k]
        ans = max(ans, curr)
        F = curr
    }

    return ans
}
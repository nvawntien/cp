func constructTransformedArray(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    rev := make([]int, n)

	for i := 0; i < n; i++ {
		rev[i] = nums[n-1-i]
	}

    for i := range nums {
        if nums[i] > 0 {
            result[i] = nums[(i + nums[i]) % n]
        } else if nums[i] < 0 {
            result[i] = rev[(n-1-i-nums[i])%n]
        } else {
            result[i] = nums[i]
        }
    }

    return result
}
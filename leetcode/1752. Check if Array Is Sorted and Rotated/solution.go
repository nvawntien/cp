func check(nums []int) bool {
    n := len(nums)
    k := n

    if nums[n-1] <= nums[0] {
        for i := 1; i < n; i++ {
            if nums[i] < nums[i-1] {
                k = i
                break
            }
        }
    }

    for i := 1; i < k; i++ {
        if nums[i] < nums[i-1] {
            return false
        }
    }

    for i := k+1; i < n; i++ {
        if nums[i] < nums[i-1] {
            return false
        }
    }

    return true
}
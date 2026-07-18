func findGCD(nums []int) int {
    MIN, MAX := nums[0], nums[0]
    for i := range nums {
        MIN = min(MIN, nums[i])
        MAX = max(MAX, nums[i])
    }

    for MAX != 0 {
        MIN, MAX = MAX, MIN % MAX
    }

    return MIN
}
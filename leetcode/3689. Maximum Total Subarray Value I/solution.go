func maxTotalValue(nums []int, k int) int64 {
    maximum := 0
    minimum := math.MaxInt

    for i := range nums {
        maximum = max(maximum, nums[i])
        minimum = min(minimum, nums[i])
    }

    return int64((maximum - minimum) * k)
}
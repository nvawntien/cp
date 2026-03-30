func minBitwiseArray(nums []int) []int {
    for i, val := range nums {
        pos := ((val + 1) & ^val) >> 1
        if pos == 0 {
            nums[i] = -1
        } else {
            nums[i] = val ^ pos
        }
    }

    return nums
}
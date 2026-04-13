func getMinDistance(nums []int, target int, start int) int {
    ans := 1_000_000_000
    abs := func (x int) int {
        mask := x >> 31
        return (x ^ mask) - mask
    }

    for i := range nums {
        if nums[i] == target && abs(i - start) < ans {
            ans = abs(i - start)
        }
    }

    return ans
}
func countMajoritySubarrays(nums []int, target int) int {
    n := len(nums)
    offset := n+1
    length := 2*n+2
    bit := make([]int, length)

    update := func (x int) {
        for ; x < length; x += x&-x {
            bit[x]++
        }
    }

    query := func (x int) int {
        res := 0
        for ; x > 0; x -= x&-x {
            res += bit[x]
        }
        return res
    }

    update(0 + offset)

    pre := 0
    ans := 0
    for i := 0; i < n; i++ {
        if nums[i] == target {
            pre++
        } else {
            pre--
        }

        ans += query(pre-1+offset)
        update(pre+offset)
    }

    return ans
}
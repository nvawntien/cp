func countOppositeParity(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)

    for i := 0; i < n; i++ {
        mod := nums[i] & 1
        cnt := 0
        for j := i+1; j < n; j++ {
            if ((nums[j] & 1) ^ mod) == 1 {
                cnt++
            }
        }

        ans[i] = cnt
    }

    return ans
}
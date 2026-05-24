func minimumSwaps(nums []int) int {
    n := len(nums)
    cnt := 0
    
    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            cnt++
        }
    }
    
    i := n-1
    ans := 0
    for cnt != 0 {
        if nums[i] != 0 {
            ans++
        }
        i--
        cnt--
    }

    return ans
}
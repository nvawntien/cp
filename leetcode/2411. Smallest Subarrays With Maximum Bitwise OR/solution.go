func smallestSubarrays(nums []int) []int {
    n := len(nums)  
    res := make([]int, n)
    last := make([]int, 30)

    for i := 0; i < n; i++ {
        res[i] = 1
    }
    
    for i := n-1; i >= 0; i-- {
        for bit := 0; bit < 30; bit++ {
            if ((nums[i] >> bit) & 1) == 1 {
                last[bit] = i
            }

            res[i] = max(res[i], last[bit] - i + 1)
        }
    }

    return res 
}
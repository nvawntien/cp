func minOperations(nums []int, k int) int {
    n := len(nums)
    ans := math.MaxInt32
    
    for x := 0; x < k; x++ {
        for y := 0; y < k; y++ {
            curr := 0
            if x == y {
                continue
            }

            for i := 0; i < n; i++ {
                if i%2==1 {
                    if nums[i] % k != y {
                        mod := nums[i] % k
                        curr += min((y-mod+k)%k, (mod-y+k)%k)
                    }
                } else {
                    if nums[i] % k != x {
                        mod := nums[i] % k
                        curr += min((x-mod+k) % k, (mod-x+k)%k)
                    }
                }
             
            }
            ans = min(ans, curr)
        }
    }

    return ans
}
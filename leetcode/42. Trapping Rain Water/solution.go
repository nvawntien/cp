func trap(height []int) int {
   // 1 0 2 
    n := len(height)
    ans := 0

    suf := make([]int, n)
    suf[n-1] = height[n-1]

    for i := n-2; i >= 0; i-- {
        suf[i] = max(suf[i+1], height[i])
    }

    l := 0
    for l < n-1 {
        if suf[l+1] < height[l] {
            height[l] = suf[l+1]
        } 

        rain := 0

        for r := l+1; r < n; r++ {
            if height[r] >= height[l] {
                l = r
                ans += rain
                break
            }
            rain += height[l] - height[r] 
        }
    }

    return ans
}
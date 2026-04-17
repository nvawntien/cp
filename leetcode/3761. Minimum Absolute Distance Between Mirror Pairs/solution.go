func minMirrorPairDistance(nums []int) int {
    pos := make(map[int]int)  
    ans := math.MaxInt32

    rev := func (x int) int {
        res := 0
        for x > 0 {
            res = res * 10 + x % 10
            x = x / 10
        }
        return res
    }

    for j, v := range nums {
        r := rev(v)
        if i, ok := pos[v]; ok {
            ans = min(ans, j - i)
        }
        pos[r] = j
    }

    if ans == math.MaxInt32 {
        return -1
    }

    return ans
}
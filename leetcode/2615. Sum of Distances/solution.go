func distance(nums []int) []int64 {
    pos := make(map[int][]int)
    n := len(nums)

    for i, x := range nums {
        pos[x] = append(pos[x], i)    
    }

    ans := make([]int64, n)

    for _, arr := range pos {
        total := 0
        for _, x := range arr {
            total += x
        }

        leftSum := 0
        m := len(arr)

        for i, x := range arr {
            rightSum := total - leftSum - x
            ans[x] = int64(i * x - leftSum + rightSum - (m-i-1) * x)
            leftSum += x
        }
    }

    return ans
}
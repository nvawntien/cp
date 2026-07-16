func gcdSum(nums []int) int64 {
    n := len(nums)
    prefixGcd := make([]int, n)
    mx := nums[0]

    gcd := func (a, b int) int {
        for b != 0 {
            a, b = b, a%b
        }
        return a
    }

    for i := 0; i < n; i++ {
        mx = max(mx, nums[i])
        prefixGcd[i] = gcd(nums[i], mx)
    }

    sort.Ints(prefixGcd)

    var sum int64 = 0
    for i := 0; i < (n>>1); i++ {
        sum += int64(gcd(prefixGcd[i], prefixGcd[n-i-1]))
    }

    return sum
}
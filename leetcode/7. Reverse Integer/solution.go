func reverse(x int) int {
    check := false
    if x < 0 {
        check = true
        x = -x
    }

    ans := 0
    for x > 0 {
        ans = ans * 10 + x % 10
        x = x / 10
    }

    if ans < math.MinInt32 || ans > math.MaxInt32 {
        return 0
    }

    if check {
        ans = -ans
    }

    return ans
}
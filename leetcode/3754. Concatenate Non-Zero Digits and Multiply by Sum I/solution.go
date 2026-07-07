func sumAndMultiply(n int) int64 {
    ans := 0
    digit := []int{}
    sum := 0
    for n > 0 {
        sum += n % 10
        if n % 10 != 0 {
            digit = append(digit, n % 10)
        }
        n /= 10
    }

    for i := len(digit)-1; i >= 0; i-- {
        ans = ans * 10 + digit[i]
    }

    return int64(ans * sum)
}
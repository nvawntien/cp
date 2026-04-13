func tribonacci(n int) int {
    f := make([]int, 40)
    f[0] = 0
    f[1] = 1
    f[2] = 1

    for i := 3; i <= n; i++ {
        f[i] = f[i-1] + f[i-2] + f[i-3]
    }

    return f[n]
}
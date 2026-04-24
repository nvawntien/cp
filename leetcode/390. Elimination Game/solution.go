func lastRemaining(n int) int {
    // f(n) is final answer if remove left to right | [1:n]
    // g(n) is final answer if remove right to left | [1:n]
    // after remove left to right, arr -> [2, 4, 6, ....n]
    // if we divide arr by 2 arr = [1, 2, 3....n/2] -> continue remove right to left with g(n/2)
    // so we can see f(n) = 2 * g(n/2)
    // easy to see f(n) + g(n) = n+1 => g(n) = n+1-f(n)
    // f(n) = 2 * [n/2 + 1 - f(n/2)]

    if n == 1 {
        return 1
    }

    return 2 * (n/2 + 1 - lastRemaining(n/2))
}
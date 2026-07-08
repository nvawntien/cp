const mod = 1_000_000_007

func sumAndMultiply(s string, queries [][]int) []int {
    n := len(s)
    left := make([]int, n)
    right := make([]int, n)
    digit := []int{0}

    for i := 0; i < n; i++ {
        d := int(s[i] - '0')
        if d != 0 {
            digit = append(digit, d)
        }
        left[i] = len(digit)-1
    }

    next := len(digit)
    for i := n-1; i >= 0; i-- {
        d := int(s[i] - '0')
        if d != 0 {
            next = left[i]
        }
        right[i] = next
    }

    m := len(digit)
    pre := make([]int, m)
    pow := make([]int, m)
    num := make([]int, m)
    pow[0] = 1

    for i := 1; i < m; i++ {
        pre[i] = (pre[i-1] + digit[i]) % mod
        num[i] = (num[i-1] * 10 + digit[i]) % mod
        pow[i] = pow[i-1] * 10 % mod
    }

    ans := make([]int, len(queries))
    for i, q := range queries {
        l, r := right[q[0]], left[q[1]]
        if l > r {
            ans[i] = 0
        } else {
            ans[i] = ((num[r] - num[l-1] * pow[r-l+1] + mod * mod) % mod) * (pre[r] - pre[l-1]) % mod
        }
    }

    return ans
}
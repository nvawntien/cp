const MOD = 1000000007

func zigZagArrays(n int, l int, r int) int {
    if n == 1 {
        return (r - l + 1) % MOD
    }

    m := r - l + 1
    V := make([]int, 2*m)

    for y := 1; y <= m; y++ {
        V[y-1] = y - 1      
        V[y-1+m] = m - y   
    }
    
    if n == 2 {
        sum := 0
        for _, v := range V {
            sum = (sum + v) % MOD
        }
        return sum
    }
    
    T := make([][]int, 2*m)
    for i := range T {
        T[i] = make([]int, 2*m)
    }
    
    for u := 1; u <= m; u++ {
        for v := 1; v < u; v++ {
            T[v-1+m][u-1] = 1
        }
        for v := u + 1; v <= m; v++ {
            T[v-1][u-1+m] = 1
        }
    }
    
    M := power(T, n-2)
    
    ans := 0
    for i := 0; i < 2*m; i++ {
        val := 0
        for j := 0; j < 2*m; j++ {
            val = (val + M[i][j]*V[j]) % MOD
        }
        ans = (ans + val) % MOD
    }
    return ans
}

func power(mat [][]int, p int) [][]int {
    n := len(mat)
    res := make([][]int, n)
    for i := range res {
        res[i] = make([]int, n)
        res[i][i] = 1
    }
    base := mat
    for p > 0 {
        if p%2 == 1 {
            res = multiply(res, base)
        }
        base = multiply(base, base)
        p /= 2
    }
    return res
}

func multiply(a, b [][]int) [][]int {
    n := len(a)
    c := make([][]int, n)
    for i := range c {
        c[i] = make([]int, n)
    }

    for i := 0; i < n; i++ {
        for k := 0; k < n; k++ {
            if a[i][k] == 0 {
                continue
            }
            aik := a[i][k]
            for j := 0; j < n; j++ {
                c[i][j] = (c[i][j] + aik*b[k][j]) % MOD
            }
        }
    }
    return c
}
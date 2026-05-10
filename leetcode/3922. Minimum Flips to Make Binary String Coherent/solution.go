func minFlips(s string) int {
    a, b := 0, 0
    n := len(s)
    
    for _, bit := range s {
        if bit == '1' {
            a++
        } else {
            b++
        }
    }

    ans := a
    if a > 0 && a-1 < ans {
        ans = a-1
    }

    if b < ans {
        ans = b
    }

    end := 0
    if s[0] == '0' {
        end++
    }
    if s[n-1] == '0' {
        end++
    }

    for i := 1; i < n-1; i++ {
        if s[i] == '1' {
            end++
        }
    }

    if end < ans {
        ans = end
    }

    return ans
}
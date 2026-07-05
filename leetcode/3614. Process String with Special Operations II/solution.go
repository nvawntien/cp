func processStr(s string, k int64) byte {
    m := int64(0)
    for i := range s {
        c := s[i]
        if c == '*' {
            if m > 0 {
                m--
            }
        } else if c == '#' {
            m = m << 1
        } else if c == '%' {
            continue
        } else {
            m++
        }
    }
    
    if k >= m {
        return '.'
    }

    for i := len(s)-1; i >= 0; i-- {
        c := s[i]
        if c == '*' {
            m++
        } else if c == '#' {
            m >>= 1
            if k >= m {
                k -= m
            }
        } else if c == '%' {
            k = m - 1 -k
        } else {
            m--
            if k == m {
                return c
            }
        }
    }

    return '.'
}
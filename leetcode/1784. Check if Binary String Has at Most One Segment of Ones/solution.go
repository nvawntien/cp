func checkOnesSegment(s string) bool {
    n := len(s)
    first := -1
    last := -1

    for i := 0; i < n; i++ {
        if s[i] == '1' {
            first = i
            break
        }
    }

    last = first+1

    for i := first+1; i < n; i++ {
        if s[i] == '1' {
            last = i
            break
        }
    }
    
    if last - first > 1 {
        return false
    }

    for i := last+1; i < n; i++ {
        if s[i] == '1' {
            first = last
            last = i
            if last - first > 1 {
                return false
            }
        }
    }

    return true
}
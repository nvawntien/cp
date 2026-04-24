func furthestDistanceFromOrigin(moves string) int {
    l, r, m := 0, 0, 0
    for _, c := range moves {
        switch c {
            case 'L':
                l++
            case 'R':
                r++
            default:
                m++
        }
    }

    if l > r {
        return l-r+m
    } 
    return r-l+m
}
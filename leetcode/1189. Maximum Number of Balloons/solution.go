func maxNumberOfBalloons(text string) int {
    b, a, l, o, n := 0, 0, 0, 0, 0

    for i := range text {
        c := text[i]
        if c == 'b' {
            b++
        } else if c == 'a' {
            a++
        } else if c == 'l' {
            l++
        } else if c == 'o' {
            o++
        } else if c == 'n' {
            n++
        }
    }

    return min(min(b, a), min(min(l/2, o/2), n))
}
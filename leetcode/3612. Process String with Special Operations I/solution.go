func processStr(s string) string {
    var buf []byte

    for i := 0; i < len(s); i++ {
        c := s[i]
        if c == '*' {
            if len(buf) > 0 {
                buf = buf[:len(buf)-1]
            }
        } else if c == '#' {
            buf = append(buf, buf...)
        } else if c == '%' {
            for l, r := 0, len(buf)-1; l < r; l, r = l+1, r-1 {
                buf[l], buf[r] = buf[r], buf[l]
            }
        } else {
            buf = append(buf, c)
        }
    }

    return string(buf)
}
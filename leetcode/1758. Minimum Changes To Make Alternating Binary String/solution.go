func minOperations(s string) int {
    zero := 0
    one := 1

    a, b := 0, 0
    for i := range s {
        c := int(s[i]-'0')
        a += c ^ zero
        b += c ^ one
        zero = 1 - zero
        one = 1 - one
    }

    return min(a, b)
}
func addBinary(a string, b string) string {
    for len(a) < len(b) {
        a = "0" + a
    }

    for len(b) < len(a) {
        b = "0" + b
    }

    n := len(a)
    res := ""
    carry := 0

    for i := n-1; i >= 0; i-- {
            sum := int(a[i] - '0') + int(b[i] - '0') + carry
            res = string(byte((sum & 1) + '0')) + res
            carry = sum >> 1
    }

    if carry > 0 {
        res = string(byte(carry + '0')) + res
    }

    return res
}

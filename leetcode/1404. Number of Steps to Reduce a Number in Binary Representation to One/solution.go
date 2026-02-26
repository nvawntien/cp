func numSteps(s string) int {
    steps := 0
    carry := 0
    
    for  i := len(s)-1; i > 0; i-- {
        bit := int(s[i]-'0') + carry
        if bit & 1 == 0 {
            steps++
        } else {
            steps += 2
            carry = 1
        }
    }

    return steps + carry
}
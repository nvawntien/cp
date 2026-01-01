func plusOne(digits []int) []int {
    carry := 0
    n := len(digits)
    
    sum := digits[n-1] + 1
    digits[n-1] = sum % 10
    carry = sum / 10
    for i := n-2; i >= 0; i-- {
        digits[i] = digits[i] + carry
        carry = digits[i] / 10
        digits[i] = digits[i] % 10
    }

    if carry != 0 {
        return append([]int{carry}, digits...)
    }

    return digits
}
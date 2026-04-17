func bitwiseComplement(n int) int {
    ans := 0
    pow := 1
    if n == 0 {
        return 1
    }
    
    for n > 0 {
        if n&1 == 0 {
            ans += pow 
        }

        pow <<= 1
        n >>= 1
    }

    return ans
}
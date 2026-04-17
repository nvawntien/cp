func binaryGap(n int) int {
    i := 0
    for n & 1 == 0 {
        i++
        n >>= 1
    }

    n >>= 1
    cur := i+1
    ans := 0

    for n > 0 {
        if n & 1 == 1 {
           if cur-i > ans {
                ans = cur-i
           }        
           i = cur
        }

        cur++
        n >>= 1
    }

    return ans
}
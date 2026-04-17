func longestPalindrome(s string) string {
    n := len(s)
    palin := make([][]bool, n)

    for i := 0; i < n; i++ {
        palin[i] = make([]bool, n)
        palin[i][i] = true
    }

    start, maxLength := 0, 1

    for length := 2; length <= n; length++ {
        for l := 0; l <= n - length; l++ {
            r := l + length - 1
            if s[l] == s[r] {
                if length == 2 {
                    palin[l][r] = true
                } else {
                    palin[l][r] = palin[l+1][r-1]
                }
            }

            if palin[l][r] && length >= maxLength {
                maxLength = length
                start = l
            }
        }
    }

    return s[start : start + maxLength]
}
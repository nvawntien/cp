func longestBalanced(s string) int {
    cnt := 1
    n := len(s)

    for l := 0; l < n; l++ {
        freq := make([]int, 26)
        uniq, maxF, cntMax := 0, 0, 0
        for r := l; r < n; r++ {
            freq[s[r]-'a']++
            f := freq[s[r] - 'a']
            
            if f == 1 {
                uniq++
            }
             
            if f > maxF {
                cntMax = 1
                maxF = f
            } else if f == maxF {
                cntMax++
            }

            if uniq == cntMax {
                if r-l+1 > cnt {
                    cnt = r-l+1
                }
            }
        }
    }

    return cnt
}
const base = 31
const mod = 1_000_000_007

func numOfStrings(patterns []string, word string) int {
    n := len(word)
    pow := make([]int, n+1)
    hashW := make([]int, n+1)

    pow[0] = 1
    for i := 1; i <= n; i++ {
        pow[i] = (pow[i-1] * base) % mod
    }

    word = " " + word
    for i := 1; i <= n; i++ {
        hashW[i] = (hashW[i-1]*base + int(word[i]-'a' + 1)) % mod
    }

    get := func(i, j int) int {
        res := (hashW[j] - hashW[i-1]*pow[j-i+1]) % mod
        if res < 0 {
            res += mod
        }
        return res
    }

    cnt := 0

    for _, pattern := range patterns {
        pLen := len(pattern)
        if pLen > n {
            continue
        }

        pHash := 0
        for i := 0; i < pLen; i++ {
            pHash = (pHash*base + int(pattern[i]-'a' + 1)) % mod
        }

        for j := 1; j <= n-pLen+1; j++ {
            if pHash == get(j, j+pLen-1) {
                cnt++
                break
            }
        }
    }

    return cnt
}
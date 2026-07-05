func numberOfSubstrings(s string) int {
    ans, l, n := 0, 0, len(s)
    freq := [3]int{}

    for r := 0; r < n; r++ {
        freq[s[r]-'a']++
        for freq[0] > 0 && freq[1] > 0 && freq[2] > 0 {
            ans += n - r 
            freq[s[l]-'a']--
            l++
        }
    }
    return ans
}
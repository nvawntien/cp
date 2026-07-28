func smallestPalindrome(s string) string {
    freq := make(map[rune]int)

	for i := range s {
		freq[rune(s[i])]++
	}

	mid := ""
	
	var leftBuf strings.Builder

	for c := 'a'; c <= 'z'; c++ {
		if freq[c] & 1 == 1 {
			mid = string(c)
			freq[c]--
		}
		n := freq[c]
		for i := 0; i < n / 2; i++ {
			leftBuf.WriteRune(c)
			freq[c]--
		}
	}

	var rightBuf strings.Builder

	for c := 'z'; c >= 'a'; c-- {
		for i := 0; i < freq[c]; i++ {
			rightBuf.WriteRune(c)
		}
	}

	return leftBuf.String() + mid + rightBuf.String()
}
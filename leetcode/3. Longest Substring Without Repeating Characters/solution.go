func lengthOfLongestSubstring(s string) int {
    lastIndex := make(map[byte]int)
    maxLength := 0
    start := 0

    for i := 0; i < len(s); i++ {
        if idx, ok := lastIndex[s[i]]; ok && idx >= start {
            start = idx + 1
        } 
        lastIndex[s[i]] = i
        maxLength = max(maxLength, i - start + 1)
    }

    return maxLength
}
func getLongestSubsequence(words []string, groups []int) []string {
    cur := groups[0]
    result := []string {words[0]}

    for i := 1; i < len(groups); i++ {
        if cur != groups[i] {
            result = append(result, words[i])
            cur = groups[i]
        }
    }

    return result
}
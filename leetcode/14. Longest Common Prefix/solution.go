func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    minLength := len(strs[0])
    prefix := ""

    for _, str := range strs {
        minLength = min(minLength, len(str))
    }

    for i := 0; i < minLength; i++ {
        check := false
        for _, str := range strs {
            if str[i] != strs[0][i] {
                check = true
                break
            }
        }
        if check {
            break
        } else {
            prefix = strs[0][:i+1]
        } 
    }

    return prefix
}
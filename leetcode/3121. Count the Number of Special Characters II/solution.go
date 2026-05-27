func numberOfSpecialChars(word string) int {
    upper := make([]int, 26)
    lower := make([]int, 26)

    for i := 0; i < 26; i++ {
        lower[i] = -1
        upper[i] = -1
    }

    for i := 0; i < len(word); i++ {
        if word[i] >= 'a' && word[i] <= 'z' {
            lower[word[i]-'a'] = i
        } else {
            if upper[word[i]-'A'] != -1 {
                continue
            }        
            upper[word[i]-'A'] = i
        }
    }

    cnt := 0

    for c := byte('a'); c <= 'z'; c++ {
        i := lower[c-'a']
        j := upper[(c^32)-'A']
        
        if i != -1 && j != -1 && i < j {
            cnt++
        }
    }

    return cnt
}
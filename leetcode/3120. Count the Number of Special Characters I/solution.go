func numberOfSpecialChars(word string) int {
    exist := make(map[byte]bool)
    for i := 0; i < len(word); i++ {
        exist[word[i]] = true
    }

    cnt := 0

    for c := byte('a'); c <= 'z'; c++ {
        if exist[c] && exist[c^32] {
            cnt++
        }
    }

    return cnt
}
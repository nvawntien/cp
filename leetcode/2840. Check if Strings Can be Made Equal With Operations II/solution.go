func checkStrings(s1 string, s2 string) bool {
    count := make([][]int, 2)

    for i := range count {
        count[i] = make([]int, 26)
    }

    for i := 0; i < len(s1); i++ {
        count[i%2][int(s1[i]-'a')]++
        count[i%2][int(s2[i]-'a')]--
    }

    for i := 0; i < len(s1); i++ {
        if count[i%2][int(s1[i]-'a')] != 0 {
            return false
        }
    }

    return true
}
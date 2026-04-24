func findTheDifference(s string, t string) byte {
    freq := make(map[byte]int)
    for _, c := range s {
        freq[byte(c)]++
    }

    for _, c := range t {
        if freq[byte(c)] == 0 {
            return byte(c)
        } 

        freq[byte(c)]--
    }

    return '2'
}
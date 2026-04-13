func maxFreqSum(s string) int {
    vowel := make(map[rune]int)
    consonants := make(map[rune]int)

    for _, c := range s {
        if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
            vowel[c]++
        } else {
            consonants[c]++
        }
    }

    freq := 0
    ans := 0
    for _, val := range vowel {
        if val > freq {
            freq = val
        }
    }

    ans += freq
    freq = 0

    for _, val := range consonants {
        if val > freq {
            freq = val
        }
    }

    ans += freq

    return ans
}

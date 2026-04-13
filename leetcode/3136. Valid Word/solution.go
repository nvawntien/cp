func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }

    vowel, consonant, dif := false, false, false

    for _, e := range word {
        c := unicode.ToLower(e)
        if c >= '0' && c <= '9' {
            continue
        } else if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c =='u'{
            vowel = true
        } else if c >= 'a' && c <= 'z' {
            consonant = true
        }else {
            dif = true
        }

    }

    return vowel && consonant && !dif
}
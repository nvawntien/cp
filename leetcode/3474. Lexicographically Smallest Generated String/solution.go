func generateString(S string, T string) string {
    n, m := len(S), len(T)
    word := make([]byte, n+m-1)
    fixed := make([]bool, n+m-1)

    for i := 0; i < n+m-1; i++ {
        word[i] = 'a'
    }

    for i := 0; i < n; i++ {
        if S[i] == 'T' {
            for j := i; j < i+m; j++ {
                if fixed[j] && word[j] != T[j-i]{
                    return ""
                } 
                
                word[j] = T[j-i]
                fixed[j] = true
            }
        }    
    }

    for i := 0; i < n; i++ {
        if S[i] == 'F' {
            match := true
            for j := i; j < i+m; j++ {
                if word[j] != T[j-i] {
                    match = false
                    break
                }
            }

            if match {
                flipped := false

                for j := i+m-1; j >= i; j-- {
                    if !fixed[j] {
                        word[j] = 'b'
                        fixed[j] = true
                        flipped = true
                        break
                    }
                }

                if !flipped {
                    return ""
                }
            }
        }
    }

    return string(word)
}
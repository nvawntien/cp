func closestTarget(words []string, target string, startIndex int) int {
    n := len(words)
    idx := -1
    for i := range words {
        if words[i] == target {
            idx = i
        }
    }

    if idx == -1 {
        return -1
    }

    i := startIndex
    right := 0
    left := 0

    for words[i] != target {
        right++
        i = (i + 1) % n 
    }

    i = startIndex

    for words[i] != target {
        left++
        i = (i-1+n) % n
    }

    return min(left, right)
}
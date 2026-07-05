func maximumLength(nums []int) int {
    counts := make(map[int64]int)
    for _, num := range nums {
        counts[int64(num)]++
    }

    maxLen := 1

    if count, exists := counts[1]; exists {
        if count%2 == 0 {
            maxLen = max(maxLen, count-1)
        } else {
            maxLen = max(maxLen, count)
        }
        delete(counts, 1)
    }

    for x := range counts {
        currLen := 0
        currX := x

        for counts[currX] >= 2 {
            currLen += 2
            currX = currX * currX
        }

        if counts[currX] >= 1 {
            currLen += 1 
        } else {
            currLen -= 1
        }

        if currLen > maxLen {
            maxLen = currLen
        }
    }

    return maxLen
}

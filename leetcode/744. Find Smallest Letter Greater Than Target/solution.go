func nextGreatestLetter(letters []byte, target byte) byte {
    ans := letters[0]

    for _, c := range letters {
        if c > target {
            return c
        }
    } 

    return ans
}
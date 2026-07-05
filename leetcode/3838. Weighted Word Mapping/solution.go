func mapWordWeights(words []string, weights []int) string {
    var ans string = ""
    for i := range words {
        wei := 0
        for j := range words[i] {
            wei += weights[int(words[i][j] - 'a')]
        }
        ans += string(byte(26 - (wei % 26) - 1) + 'a')
    }

    return ans
}
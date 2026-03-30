func largestSubmatrix(matrix [][]int) int {
    ans := 0
    
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] != 0 && i > 0 {
                matrix[i][j] += matrix[i-1][j]
            }
        }

        col := make([]int, len(matrix[i]))
        copy(col, matrix[i])

        sort.Slice(col, func (i, j int) bool {
            return col[i] > col[j]
        })

        for k := range col {
            ans = max(ans, col[k] * (k+1))
        }
    }

    return ans
}
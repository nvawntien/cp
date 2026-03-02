func minSwaps(grid [][]int) int {
    n := len(grid)
    zero := make([]int, n)

    for i := 0; i < n; i++ {
        count := 0
        for j := n-1; j >= 0 && grid[i][j] == 0; j-- {
            count++
        }
        zero[i] = count
    }

    step := 0

    for i := 0; i < n; i++ {
        need := n-i-1
        j := i
        for j < n && zero[j] < need {
            j++
        }

        if j == n {
            return -1
        }

        for j > i {
            temp := zero[j-1]
            zero[j-1] = zero[j]
            zero[j] = temp
            step++
            j--
        }
    }

    return step
}
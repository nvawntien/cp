func minOperations(grid [][]int, x int) int {
    n := len(grid)
    m := len(grid[0])
    N := n * m
    arr := make([]int, n*m)

    for i := range grid  {
        for j := range grid[i] {
            arr[i*m+j] = grid[i][j]
        }
    }

    sort.Ints(arr)

    diff := 0
    abs := func (x int) int {
        mask := x >> 31
        return (x ^ mask) - mask
    }

    tar := arr[N / 2]
    for i := range arr {
        curr := abs(arr[i] - tar)
        if curr % x != 0 {
            return -1
        }
        diff += curr / x
    }

    return diff
}
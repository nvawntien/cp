func maxProduct(n int) int {
    arr := []int {}

    for n > 0 {
        arr = append(arr, n%10)
        n /= 10
    }

    sort.Slice(arr, func (i, j int) bool {
        return arr[i] > arr[j]
    })

    return arr[0] * arr[1]
}
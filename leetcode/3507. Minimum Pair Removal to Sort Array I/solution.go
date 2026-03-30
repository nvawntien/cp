func minimumPairRemoval(nums []int) int {
    var ops = 0
    arr := append([]int{}, nums...)

    for !isSorted(arr) {
        minSum := int(1e9)
        idx := 0
        for i := 0; i < len(arr)-1; i++ {
            if arr[i] + arr[i+1] < minSum {
                minSum = arr[i] + arr[i+1]
                idx = i
            }
        }

        arr[idx] = minSum
        arr = append(arr[:idx+1], arr[idx+2:]...)
        ops++
    }

    return ops
}

func isSorted(arr []int) bool {
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] > arr[i+1] {
            return false
        }
    }

    return true
}
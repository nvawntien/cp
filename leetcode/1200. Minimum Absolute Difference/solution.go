func minimumAbsDifference(arr []int) [][]int {
    sort.Slice(arr, func (i, j int) bool {
        return arr[i] < arr[j]
    })

    ans := [][]int{}
    absMin := int(1e9)

    for i := 0; i < len(arr)-1; i++ {
        absMin = min(absMin, abs(arr[i] - arr[i+1]))
    }

    for i := 0; i < len(arr)-1; i++ {
        if abs(arr[i] - arr[i+1]) == absMin {
            ans = append(ans, []int{arr[i], arr[i+1]})
        }
    }

    return ans
}

func abs(a int) int {
    if a < 0 {
        return -a
    }

    return a
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}
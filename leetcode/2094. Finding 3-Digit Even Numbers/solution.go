func findEvenNumbers(digits []int) []int {
    arr := []int {}
    n := len(digits)
    freq := make(map[int]bool)

    for i := 0; i < n; i++ {
        if digits[i] == 0 {
            continue
        }

        for j := 0; j < n; j++ {
            if i == j {
                continue
            }

            for k := 0; k < n; k++ {

                if digits[k] % 2 == 1 || k == j || k == i {
                    continue
                }

                num := digits[i] * 100 + digits[j] * 10 + digits[k]
                if _, exists := freq[num]; !exists {
                    freq[num] = true
                    arr = append(arr, num)
                }
            }
        }
    }

    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })

    return arr
}
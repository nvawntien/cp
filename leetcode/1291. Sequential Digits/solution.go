func sequentialDigits(low int, high int) []int {
    ans := []int{}

    for size := 2; size <= 9; size++ {
        for start := 1; start <= 10 - size; start++ {
            num := 0
            for i := 0; i < size; i++ {
                num = num * 10 + start + i
            }

            if num >= low && num <= high {
                ans = append(ans, num)
            }
        }
    }

    return ans
}
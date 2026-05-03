func rotatedDigits(n int) int {
    ans := 0

    for i := 1; i <= n; i++ {
        valid := true
        diff := false
        num := i
        for num > 0 {
            digit := num % 10
            if digit == 3 || digit == 4 || digit == 7 {
                valid = false
                break
            }

            if digit == 2 || digit == 5 || digit == 6 || digit == 9 {
                diff = true
            }

            num /= 10
        }

        if valid && diff {
            ans++
        }
    }

    return ans
}
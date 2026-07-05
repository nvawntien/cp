func totalWaviness(num1 int, num2 int) int {
    ans := 0
    for i := max(101, num1); i <= num2; i++ {
        num := i
        for num > 0 && num / 10 > 0 && num / 100 > 0{
            r := num % 10

            c := (num / 10) % 10
            l := (num / 100) % 10

            if (c > r && c > l) || (c < r && c < l) {
                ans++
            }

            num /= 10
        }
    }

    return ans
}
func sumOfPrimesInRange(n int) int {
    r := 0
    temp := n
    for temp > 0 {
        r = r * 10 + temp % 10
        temp /= 10
    }

    sum := 0
    
    for i := min(r, n); i <= max(r, n); i++ {
        if i == 1 {
            continue
        }
        num := i
        check := true
        for j := 2; j * j <= num; j++ {
            if num % j == 0 {
                check = false
                break
            }
        }

        if check {
            sum += i
        }
    }

    return sum
}
func numOfUnplacedFruits(fruits []int, baskets []int) int {
    n := len(fruits)
    tick := make([]bool, n)

    for _, fruit := range fruits {
        for i := 0; i < n; i++ {
            if !tick[i] && baskets[i] >= fruit {
                tick[i] = true
                break
            }
        }
    }

    count := 0

    for i := 0; i < n; i++ {
        if !tick[i] {
            count++
        }
    }

    return count
}
func maxIceCream(costs []int, coins int) int {
    cost := 0
    for i := range costs {
        cost = max(cost, costs[i])
    }

    count := make([]int, cost+1)
    for i := range costs {
        count[costs[i]]++
    }

    cnt := 0
    for i := range count {
        if coins <= 0 || coins < i {
            break
        }

        for count[i] > 0 && coins >= i {
            cnt++
            coins -= i
            count[i]--
        }
    }

    return cnt
}
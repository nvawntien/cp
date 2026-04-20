func maxDistance(colors []int) int {
    n := len(colors)
    l, r := n-1, 0

    for i := range colors {
        if colors[n-1] ^ colors[i] > 0 {
            l = i
            break
        }
    }

    for i := range colors {
        if colors[0] ^ colors[i] > 0 {
            r = i
        }
    }

    return max(n-1-l, r)
}
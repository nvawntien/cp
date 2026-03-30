func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
    var maxArea int64 = 0
    n := len(bottomLeft)

    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            interBLX := max(bottomLeft[i][0], bottomLeft[j][0])
            interBLY := max(bottomLeft[i][1], bottomLeft[j][1])
            interTRX := min(topRight[i][0], topRight[j][0])
            interTRY := min(topRight[i][1], topRight[j][1])

            width := interTRX - interBLX
            height := interTRY - interBLY

            if width > 0 && height > 0 {
                side := min(width, height)
                area := int64(side) * int64(side)
                if area > maxArea {
                    maxArea = area
                }
            }
        }
    }

    return maxArea
}

func max(a, b int) int {
    if a < b {
        return b
    }

    return a
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}
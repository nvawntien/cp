func canReach(s string, minJump int, maxJump int) bool {
    n := len(s)
    if s[n-1] != '0' {
        return false
    }

    queue := make([]int, 0, n)
    queue = append(queue, 0)
    farthest := 0
    
    for len(queue) > 0 {
        i := queue[0]
        queue = queue[1:]

        if i == n-1 {
            return true
        }

        start := max(i+minJump, farthest+1)
        end := min(i+maxJump, n-1)

        for j := start; j <= end; j++ {
            if s[j] == '0' {
                queue = append(queue, j)
            }
        }

        farthest = max(farthest, end)
    }

    return false
}
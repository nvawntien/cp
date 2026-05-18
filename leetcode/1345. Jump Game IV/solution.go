func minJumps(nums []int) int {
    n := len(nums)
    graph := make(map[int][]int)

    for i, num := range nums {
        graph[num] = append(graph[num], i)
    }

    visited := make([]bool, n)
    queue := []int {0}
    visited[0] = true
    steps := 0

    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            u := queue[0]
            queue = queue[1:]

            if u == n-1 {
                return steps
            }

            if u+1 < n && !visited[u+1] {
                visited[u+1] = true
                queue = append(queue, u+1)
            }

            if u-1 >= 0 && !visited[u-1] {
                visited[u-1] = true
                queue = append(queue, u-1)
            }

            for _, idx := range graph[nums[u]] {
                if !visited[idx] {
                    visited[idx] = true
                    queue = append(queue, idx)
                }
            }

            delete(graph, nums[u])
        }

        steps++
    }

    return -1
}
func canReach(arr []int, start int) bool {
    n := len(arr)

    visited := make([]bool, n)
    queue := []int {start}
    visited[start] = true

    for len(queue) > 0{
        u := queue[0]
        queue = queue[1:]

        if arr[u] == 0 {
            return true
        }

        if u-arr[u] >= 0 && !visited[u-arr[u]] {
            queue = append(queue, u-arr[u])
            visited[u-arr[u]] = true
        }

        if u+arr[u] < n && !visited[u+arr[u]] {
            queue = append(queue, u+arr[u])
            visited[u+arr[u]] = true
        }
    } 

    return false
}
func minJumps(nums []int) int {
    n := len(nums)
    maximum := 0

    for _, num := range nums {
        maximum = max(maximum, num)
    }

    spf := make([]int, maximum+1)
    for i := range spf {
        spf[i] = i
    }

    for i := 2; i <= maximum; i++ {
        if spf[i] == i {
            for j := i + i; j <= maximum; j+=i {
                spf[j] = i
            }
        }
    }

    mul := make(map[int][]int)

    for i, num := range nums {
        curr := num
        for curr > 1 {
            prime := spf[curr]
            mul[prime] = append(mul[prime], i)
            for curr%prime == 0 {
                curr /= prime
            }
        }
    }

    visited := make([]bool, n)
    queue := []int{0}
    visited[0] = true

    step := make([]int, n)

    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]

        if u == n-1 {
            return step[u]
        }

        if u-1 >= 0 && !visited[u-1] {
            queue = append(queue, u-1)
            visited[u-1] = true
            step[u-1] = step[u] + 1
        }

        if u+1 < n && !visited[u+1] {
            queue = append(queue, u+1)
            visited[u+1] = true
            step[u+1] = step[u] + 1
        }

        if nums[u] >= 2 && spf[nums[u]] == nums[u] {
            for _, p := range mul[nums[u]] {
                if !visited[p] {
                    queue = append(queue, p)
                    visited[p] = true
                    step[p] = step[u] + 1
                }
            }
            delete(mul, nums[u])
        }
    }

    return step[n-1]
}
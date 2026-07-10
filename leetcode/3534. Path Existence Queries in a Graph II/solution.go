type Node struct {
    val int
    idx int
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
    nodes := make([]Node, n)
    for i := range nodes {
        nodes[i] = Node{val: nums[i], idx: i}
    }

    sort.Slice(nodes, func (i, j int) bool {
        return nodes[i].val < nodes[j].val
    })

    pos := make([]int, n)
    for i := range pos {
        pos[nodes[i].idx] = i
    }

    LOG := 1
    for (1 << LOG) <= n {
        LOG++
    }

    jump := make([][]int, n)
    for i := range jump {
        jump[i] = make([]int, LOG)
    }

    for l, r := 0, 0; l < n; l++ {
        if r < l {
            r = l
        }
        for r+1 < n && nodes[r+1].val - nodes[l].val <= maxDiff {
            r++
        }
        jump[l][0] = r
    }

    for j := 1; j < LOG; j++ {
        for i := 0; i < n; i++ {
            jump[i][j] = jump[jump[i][j-1]][j-1]
        }
    }

    ans := make([]int, len(queries))
    for i, query := range queries {
        left, right := pos[query[0]], pos[query[1]]
        if left == right {
            ans[i] = 0
            continue
        }

        if left > right {
            left, right = right, left
        }

        distance := 0
        for j := LOG-1; j >= 0; j-- {
            if jump[left][j] < right {
                left = jump[left][j]
                distance += 1 << j
            }
        }

        if jump[left][0] >= right {
            ans[i] = distance + 1
        } else {
            ans[i] = -1
        }
    }

    return ans
}
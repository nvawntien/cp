func getResults(queries [][]int) []bool {
    mx := 0
    for _, q := range queries {
        mx = max(mx, q[1])
    }

    n := mx+2

    obs := make([]bool, n)
    obs[0] = true
    obs[n-1] = true
    
    for _, q := range queries {
        if q[0] == 1 {
            obs[q[1]] = true
        }
    }

    parentL := make([]int, n)
    parentR := make([]int, n)

    for i := range parentL {
        parentL[i] = i
        parentR[i] = i
    }

    var find func(int, []int) int
    find = func (i int, parent []int) int {
        if parent[i] == i {
            return i
        }
        parent[i] = find(parent[i], parent)
        return parent[i]
    }

    for i := 1; i < n; i++ {
        if !obs[i] {
            parentL[i] = i-1
            parentR[i] = i+1
        }
    }

    bit := make([]int, n+1)

    update := func(idx, val int) {
        for i := idx+1; i < len(bit); i += i&-i {
            bit[i] = max(bit[i], val)
        }
    }

    get := func (idx int) int {
        res := 0
        for i := idx+1; i > 0; i -= i&-i {
            res = max(res, bit[i])
        }
        return res
    }

    for i := 1; i < n; i++ {
        if obs[i] {
            prev := find(i-1, parentL)
            update(i, i-prev)
        }
    }

    var ans []bool
    for i := len(queries)-1; i >= 0; i-- {
        q := queries[i]
        x := q[1]

        if q[0] == 1 {
            prev := find(x-1, parentL)
            next := find(x+1, parentR)

            parentL[x] = x-1
            parentR[x] = x+1

            update(next, next-prev)
        } else {
            sz := q[2]
            prev := find(x, parentL)
            gap := max(get(prev), x-prev)

            ans = append(ans, gap >= sz)
        }
    }

    for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

    return ans
}
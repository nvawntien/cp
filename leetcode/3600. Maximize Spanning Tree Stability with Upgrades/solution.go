type DSU struct {
    parent []int
}

func NewDSU(n int) *DSU {
    p := make([]int, n)
    for i := range p {
        p[i] = i
    }
    return &DSU{
        parent: p,
    }
}

func (d *DSU) Find(v int) int {
    if v == d.parent[v] {
        return v
    }

    d.parent[v] = d.Find(d.parent[v])
    return d.parent[v]
}

func (d *DSU) Union(u, v int) bool{
    u = d.Find(u)
    v = d.Find(v)

    if u != v {
        d.parent[v] = u
        return true
    }

    return false
}

func maxStability(n int, edges [][]int, k int) int {
    dsuTest := NewDSU(n)

    for _, edge := range edges {
        u, v, must := edge[0], edge[1], edge[3]
        if must == 1 {
            if !dsuTest.Union(u, v) {
                return -1
            }
        } 
    }

    dsuAll := NewDSU(n)
    need := n

    for _, edge := range edges {
        u, v := edge[0], edge[1]
        if dsuAll.Union(u, v) {
            need--
        }
    }

    if need > 1 {
        return -1
    }

    check := func (x int) bool {
        used := 0
        upgraded := 0
        dsu := NewDSU(n)
        
        for _, edge := range edges {
            u, v, s, must := edge[0], edge[1], edge[2], edge[3]
            if must == 1 {
                if s < x {
                    return false
                }

                if dsu.Union(u, v) {
                    used++
                }
            }
        }

        for _, edge := range edges {
            u, v, s, must := edge[0], edge[1], edge[2], edge[3]
            if must == 0 && s >= x {
                if dsu.Union(u, v) {
                    used++
                } 
            }
        }

        for _, edge := range edges {
            u, v, s, must := edge[0], edge[1], edge[2], edge[3]
            if must == 0 && s*2 >= x && upgraded < k {
                if dsu.Union(u, v) {
                    used++
                    upgraded++
                } 
            }
        }

        return used == n-1
    }

    low, high, ans := 1, int(2e5), -1
    for low <= high {
        mid := (low + high) >> 1
        if check(mid) {
            ans = mid
            low = mid+1
        } else {
            high = mid-1
        }
    }

    return ans
}

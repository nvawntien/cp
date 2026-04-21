type DSU struct {
    parent []int
    size   []int
}

func NewDSU(n int) *DSU {
    parent := make([]int, n)
    size := make([]int, n)

    for i := 0; i < n; i++ {
        parent[i] = i
        size[i] = 1
    }

    return &DSU{
        parent: parent,
        size: size,
    }
}

func (dsu *DSU) Find(x int) int {
    if dsu.parent[x] == x {
        return x
    }

    dsu.parent[x] = dsu.Find(dsu.parent[x])
    return dsu.parent[x]
}

func (dsu *DSU) Union(a, b int) {
    a = dsu.Find(a)
    b = dsu.Find(b)

    if a == b {
        return
    }

    if dsu.size[a] < dsu.size[b] {
        a, b = b, a
    }
    dsu.parent[b] = a
    dsu.size[a] += dsu.size[b]
}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
    n := len(source)
    dsu := NewDSU(n)

    for _, swap := range allowedSwaps {
        dsu.Union(swap[0], swap[1])
    }

    for i := 0; i < n; i++ {
        dsu.parent[i] = dsu.Find(i)
    }

    count := make([]map[int]int, n)

    for i := 0; i < n; i++ {
        root := dsu.parent[i]
        if count[root] == nil {
            count[root] = make(map[int]int)
        }
        count[root][source[i]]++    
    }

    ans := 0

    for i := 0; i < n; i++ {
        root := dsu.parent[i]
        if count[root][target[i]] > 0 {
            count[root][target[i]]--
        } else {
            ans++
        }
    }

    return ans
}
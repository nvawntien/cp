package main

type Fenwick struct {
    n      int
    b1, b2 []int
}

func NewFenwick(n int) *Fenwick {
    return &Fenwick {
        n: n,
        b1: make([]int, n+1),
        b2: make([]int, n+1),
    }
}

func (f *Fenwick) add(bit []int, idx, val int) {
    for i := idx; i <= f.n; i += i&-i {
        bit[i] += val
    }
}

func (f *Fenwick) update(i, val int) {
    f.add(f.b1, i, val)
    f.add(f.b2, i, (i-1) * val)
}

func (f *Fenwick) rangeAdd(l, r, val int) {
    f.update(l, val)
    f.update(r+1, -val)
}

func (f *Fenwick) query(bit []int, idx int) int {
    sum := 0
    for i := idx; i > 0; i -= i&-i {
        sum += bit[i]
    }
    return sum
}

func (f *Fenwick) prefixSum(i int) int {
    s1 := f.query(f.b1, i)
    s2 := f.query(f.b2, i)
    return i * s1 - s2
}

func (f *Fenwick) rangeSum(l, r int) int {
    return f.prefixSum(r) - f.prefixSum(l-1)
}

func sumCounts(nums []int) int {
    n := len(nums)
    fw := NewFenwick(n)
    seen := make(map[int]int)
    ans, sq := 0, 0

    for i := 0; i < n; i++ {
        pos, ok := seen[nums[i]]

        if !ok {
            pos = -1
        }

        l, r := pos+2, i+1
        old := fw.rangeSum(l, r)
        sq += old * 2 + (r - l + 1)
        fw.rangeAdd(l, r, 1)    
        ans += sq
        seen[nums[i]] = i
    }
    
    return ans
}
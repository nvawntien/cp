type Item struct {
	val int
	l   int
	r   int
}

type MaxHeap []Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxTotalValue(nums []int, k int) int64 {
    n := len(nums)
    log := bits.Len(uint(n))

    maxST := make([][]int, n)
    minST := make([][]int, n)

    for i := range maxST {
        maxST[i] = make([]int, log)
        minST[i] = make([]int, log)
        maxST[i][0] = nums[i]
        minST[i][0] = nums[i]
    }

    for j := 1; j < log; j++ {
        for i := 0; i+(1<<j) <= n; i++ {
            maxST[i][j] = max(maxST[i][j-1], maxST[i+(1<<(j-1))][j-1])
            minST[i][j] = min(minST[i][j-1], minST[i+(1<<(j-1))][j-1])
        }
    }

    query := func (l, r int) int {
        j := bits.Len(uint(r-l+1))-1
        maxVal := max(maxST[l][j], maxST[r-(1<<j)+1][j])
        minVal := min(minST[l][j], minST[r-(1<<j)+1][j])
        return maxVal - minVal
    }

    h := &MaxHeap{}
    heap.Init(h)

    for i := 0; i < n; i++ {
        heap.Push(h, Item{
            val: query(i, n-1),
            l: i,
            r: n-1,
        })
    }

    ans := int64(0)

    for i := 0; i < k; i++ {
        item := heap.Pop(h).(Item)
        ans += int64(item.val)

        if item.l < item.r {
            heap.Push(h, Item{
                val: query(item.l, item.r-1),
                l: item.l,
                r: item.r-1,
            })
        }
    }

    return ans
}
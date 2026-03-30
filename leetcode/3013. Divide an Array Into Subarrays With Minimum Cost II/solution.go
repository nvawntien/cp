func minimumCost(nums []int, k int, dist int) int64 {
    n    := len(nums)
    temp := make([]int, n)
    copy(temp, nums)
    sort.Ints(temp)
    m := 0

    for i := 0; i < n; i++ {
        if i == 0 || temp[i] != temp[i-1] {
            temp[m] = temp[i]
            m++
        }
    }

    sorted   := temp[:m]
    bitSum   := make([]int64, m+1)
    bitCount := make([]int, m+1)
    pos      := make(map[int]int)

    for i := 0; i < m; i++ {
        pos[sorted[i]] = i+1
    }

    update := func (i, v, c int) {
        for ;i <= m; i += i & -i {
            bitSum[i]   += int64(v)
            bitCount[i] += c
        }
    }

    maxJumpLength := 1
    for maxJumpLength << 1 <= m {
        maxJumpLength <<= 1
    }

    ans := int64(math.MaxInt64)

    for i := 1; i < n; i++ {
        update(pos[nums[i]], nums[i], 1)
        if i > dist+1 {
            last := nums[i-dist-1]
            update(pos[last], -last, -1)
        }  

        if i >= k-1 {
            idx, curr, val := 0, 0, int64(0)
            for p := maxJumpLength; p > 0; p >>= 1 {
                if idx+p <= m && curr + bitCount[idx+p] < k-1 {
                    idx  += p
                    curr += bitCount[idx]
                    val  += bitSum[idx]
                    
                }
            }

            if curr < k-1 {
                val += int64(k-1-curr) * int64(sorted[idx]) 
            }

            if val < ans {
                ans = val
            }
        }      
    }

    return ans + int64(nums[0])
}
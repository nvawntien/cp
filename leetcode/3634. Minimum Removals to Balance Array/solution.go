func minRemoval(nums []int, k int) int {
    if len(nums) == 1 {
        return 0
    }

    sort.Ints(nums)
    n := len(nums)
    cnt := n

    for i := range nums {
        low, high, pos := i+1, n-1, -1
        for low <= high {
            mid := (low + high) >> 1
            if nums[mid] <= nums[i] * k {
                pos = mid
                low = mid + 1
            } else {
                high = mid - 1
            }
        }

        if pos != -1 {
            cnt = min(cnt, i + n-pos-1)
        }
    }

    if cnt == n {
        return n-1
    }
    
    return cnt
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
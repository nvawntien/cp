func pivotArray(nums []int, pivot int) []int {
    n := len(nums)
    l, e, g := 0, 0, 0
    ans := make([]int, n)

    for i := 0; i < n; i++ {
        if nums[i] < pivot {
            l++
        } else if nums[i] == pivot {
            e++   
        } else {
            g++
        }
    }

    li, ei, gi := 0, l, l+e
    
    for i := 0; i < n; i++ {
        if nums[i] < pivot {
            ans[li] = nums[i]
            li++
        } else if nums[i] == pivot { 
            ans[ei] = nums[i]
            ei++
        } else {
            ans[gi] = nums[i]
            gi++
        }
    }

    return ans
}
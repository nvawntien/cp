func isTrionic(nums []int) bool {
    for p := 1; p < len(nums)-2; p++ {
        for q := p+1; q < len(nums)-1; q++ {
            if isIncreasing(nums[:p+1]) && isDecreasing(nums[p:q+1]) && isIncreasing(nums[q:]) {
                return true
            }
        }
    }
    return false    
}

func isIncreasing(nums []int) bool {
    for i := 1; i < len(nums); i++ {
        if nums[i] <= nums[i-1] {
            return false
        }
    }
    return true
}

func isDecreasing(nums []int) bool {
    for i := 1; i < len(nums); i++ {
        if nums[i] >= nums[i-1] {
            return false
        }
    }
    return true
}
func findMin(nums []int) int {
    n := len(nums)
    low, high := 0, n-1

    for low < high {
        mid := low + (high-low) >> 1
        if nums[mid] > nums[high] {
            low = mid+1
        } else {
            high = mid
        }
    }

    return nums[low]
}
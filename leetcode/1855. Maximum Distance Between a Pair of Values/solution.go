func maxDistance(nums1 []int, nums2 []int) int {
    n := len(nums1)
    m := len(nums2)

    i, j := 0, 0
    ans := 0

    for i < n && j < m {
        if i > j {
            j++
        } else if nums1[i] > nums2[j] {
            i++
        } else {
            ans = max(ans, j-i)
            j++
        }
    }

    return ans
}
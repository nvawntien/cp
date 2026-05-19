func getCommon(nums1 []int, nums2 []int) int {
    n, m := len(nums1), len(nums2)
    i, j := 0, 0

    for i < n && j < m {
        if nums1[i] < nums2[j] {
            i++
        } else if nums1[i] > nums2[j] {
            j++
        } else {
            return nums1[i]
        }
    }

    return -1
}
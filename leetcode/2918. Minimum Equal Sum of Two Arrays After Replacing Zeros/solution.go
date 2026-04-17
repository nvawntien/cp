func minSum(nums1 []int, nums2 []int) int64 {
    sum1, sum2 := 0, 0
    n1, n2 := 0, 0

    for i := 0; i < len(nums1); i++ {
        if nums1[i] == 0 {
            n1++
            continue
        }

        sum1 += nums1[i]
    }

    for i := 0; i < len(nums2); i++ {
        if nums2[i] == 0 {
            n2++
            continue
        }

        sum2 += nums2[i]
    }

    sum1 += n1
    sum2 += n2

    if sum1 == sum2 {
        return int64(sum1)
    }

    if n1 > 0 && n2 > 0 {
        return int64(max(sum1, sum2))
    }

    if n1 == 0 && n2 > 0 && sum1 > sum2{
        return int64(sum1)
    }

    if n2 == 0 && n1 > 0 && sum2 > sum1 {
        return int64(sum2)
    }

    return -1
}
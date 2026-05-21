func longestCommonPrefix(arr1 []int, arr2 []int) int {
    pre := make(map[int]bool)
    for _, num := range arr1 {
        for num > 0 {
            pre[num] = true
            num /= 10
        }
    }

    ans := 0

    for _, num := range arr2 {
        for num > 0 {
            if _, ok := pre[num]; ok && num > ans{
                ans = num
                break
            }
            num /= 10
        }
    }

    length := 0
    for ans > 0 {
        length++
        ans /=10
    }

    return length
}
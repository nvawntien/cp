func repeatedNTimes(nums []int) int {
    mp := make(map[int]int)
    for _, number := range nums {
        mp[number]++
        if mp[number] > 1 {
            return number
        }
    }

    return 0
}
func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)
    ans := make([]int, 2)

    for i := 0; i < len(nums); i++ {
        
        if mp[target - nums[i]] > 0 {
            ans[0] = mp[target - nums[i]] - 1
            ans[1] = i               
        } else {
            mp[nums[i]] = i + 1
        }
    }
    return ans
}
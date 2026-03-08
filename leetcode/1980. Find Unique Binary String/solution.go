func findDifferentBinaryString(nums []string) string {
    n := len(nums[0])
    seen := make(map[int]bool)

    for i := 0; i < len(nums); i++ {
        seen[convertToNumber(nums[i])] = true
    }

    for i := 0; i < (1 << n); i++ {
        if !seen[i] {
            return fmt.Sprintf("%0*b", n, i)
            break
        }
    }
    return ""
}

func convertToNumber(s string) int {
    ans := 0
    for i := 0; i < len(s); i++ {
        ans = ans * 2 + int(s[i]-'0')
    }

    return ans
}
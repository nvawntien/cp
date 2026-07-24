func uniqueXorTriplets(nums []int) int {
    pairs := make([]bool, 2048)
    triplets := make([]bool, 2048)

    n := len(nums)

    for j := 0; j < n; j++ {
        for k := j; k < n; k++ {
            pairs[nums[j] ^ nums[k]] = true
        }
    }

    for i := 0; i < n; i++ {
        for p := 0; p < 2048; p++ {
            if pairs[p] {
                triplets[nums[i] ^ p] = true
            }
        }
    }

    ans := 0
    for i := 0; i < 2048; i++ {
        if triplets[i] {
            ans++
        }
    }

    return ans
}
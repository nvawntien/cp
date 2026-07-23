func uniqueXorTriplets(nums []int) int {
    n := len(nums)
    if n <= 2 {
        return n
    }
    x := bits.Len(uint(n))
    return (1 << x)
}
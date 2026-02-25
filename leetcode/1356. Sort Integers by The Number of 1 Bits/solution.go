func sortByBits(arr []int) []int {
    sort.Slice(arr, func (i, j int) bool {
        a := countOne32(arr[i])
        b := countOne32(arr[j])
        if a != b {
            return a < b
        }

        return arr[i] < arr[j]
    })

    return arr
}

func countOne32(x int) int {
	x = (x & 0x55555555) + ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x & 0x0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f)
	x = (x & 0x00ff00ff) + ((x >> 8) & 0x00ff00ff)
	x = (x & 0x0000ffff) + ((x >> 16) & 0x0000ffff)
	return x
}
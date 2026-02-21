func countPrimeSetBits(left int, right int) int {
    cnt := 0
    for i := left; i <= right; i++ {
        bit := countOne32(i)
        switch bit {
            case 2, 3, 5, 7, 11, 13, 17, 19:
                cnt++
            default:
                continue
        }
    }

    return cnt
}

func countOne32(x int) int {
	x = (x & 0x55555555) + ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x & 0x0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f)
	x = (x & 0x00ff00ff) + ((x >> 8) & 0x00ff00ff)
	x = (x & 0x0000ffff) + ((x >> 16) & 0x0000ffff)
	return x
}

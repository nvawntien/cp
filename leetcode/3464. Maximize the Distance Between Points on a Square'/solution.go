func maxDistance(side int, points [][]int, k int) int {
	n := len(points)
	P := make([]int64, n)
	S := int64(side)

	for i, pt := range points {
		x, y := int64(pt[0]), int64(pt[1])
		if y == 0 {
			P[i] = x
		} else if x == S {
			P[i] = S + y
		} else if y == S {
			P[i] = 3*S - x
		} else {
			P[i] = 4*S - y
		}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i] < P[j]
	})

	P2 := make([]int64, 2*n)
	for i := 0; i < n; i++ {
		P2[i] = P[i]
		P2[i+n] = P[i] + 4*S
	}

	up := make([][17]int32, 2*n+1)

	l, r := int64(0), S
	ans := int64(0)

	for l <= r {
		m := l + (r - l)/2
		if check(m, P2, n, k, S, up) {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return int(ans)
}

func check(m int64, P2 []int64, n int, k int, S int64, up [][17]int32) bool {
	j := 0

	for i := 0; i < 2*n; i++ {
		if j <= i {
			j = i + 1
		}
		for j < 2*n && P2[j]-P2[i] < m {
			j++
		}
		up[i][0] = int32(j)
	}
	up[2*n][0] = int32(2 * n)

	for step := 1; step < 17; step++ {
		for i := 0; i <= 2*n; i++ {
			up[i][step] = up[up[i][step-1]][step-1]
		}
	}

	targetDist := 4*S - m
	jumps := k - 1 

	for i := 0; i < n; i++ {
		curr := int32(i)
		for bit := 0; bit < 17; bit++ {
			if (jumps>>bit)&1 == 1 {
				curr = up[curr][bit]
			}
		}
		
		if int(curr) < i+n && P2[curr]-P2[i] <= targetDist {
			return true
		}
	}

	return false
}
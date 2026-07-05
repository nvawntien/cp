type Result struct {
	count    int64
	waviness int64
}

func totalWaviness(num1 int64, num2 int64) int64 {
	return solve(num2) - solve(num1-1)
}

func solve(n int64) int64 {
	if n <= 0 {
		return 0
	}
	s := strconv.FormatInt(n, 10)

	var memo [20][11][11]Result
	for i := 0; i < len(s); i++ {
		for j := 0; j < 11; j++ {
			for k := 0; k < 11; k++ {
				memo[i][j][k] = Result{-1, -1} 
			}
		}
	}

	var dp func(i int, prev1 int, prev2 int, isLimit bool, isLead bool) Result
	dp = func(i int, prev1 int, prev2 int, isLimit bool, isLead bool) Result {
		if i == len(s) {
			return Result{1, 0}
		}

		if !isLimit && !isLead && memo[i][prev1][prev2].count != -1 {
			return memo[i][prev1][prev2]
		}

		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}

		var resCnt, resW int64

		for d := 0; d <= up; d++ {
			nxtLimit := isLimit && (d == up)
			nxtLead := isLead && (d == 0)

			nxtPrev1 := prev1
			nxtPrev2 := prev2

			if nxtLead {
				nxtPrev1 = 10
				nxtPrev2 = 10
			} else {
				nxtPrev1 = d
				nxtPrev2 = prev1 
			}

			w := int64(0)
			if !isLead && prev2 != 10 {
				if prev2 < prev1 && prev1 > d {
					w = 1 
				} else if prev2 > prev1 && prev1 < d {
					w = 1 
				}
			}

			sub := dp(i+1, nxtPrev1, nxtPrev2, nxtLimit, nxtLead)

			resCnt += sub.count
			resW += sub.waviness + w*sub.count
		}

		res := Result{resCnt, resW}
		
		if !isLimit && !isLead {
			memo[i][prev1][prev2] = res
		}

		return res
	}

	return dp(0, 10, 10, true, true).waviness
}
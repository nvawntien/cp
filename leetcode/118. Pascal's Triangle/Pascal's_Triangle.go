func generate(numRows int) [][]int {
	pascal := make([][]int, numRows)

	for i := 1; i <= numRows; i++ {
		pascal[i-1] = make([]int, i)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				pascal[i][j] = 1
			} else {
				pascal[i][j] = pascal[i-1][j] + pascal[i-1][j-1]
			}
		}
	}

	return pascal
}
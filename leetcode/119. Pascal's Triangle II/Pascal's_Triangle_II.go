func getRow(rowIndex int) []int {
	pascal := make([][]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		pascal[i] = make([]int, i+1)
	}

	for i := 0; i <= rowIndex; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				pascal[i][j] = 1
			} else {
				pascal[i][j] = pascal[i-1][j-1] + pascal[i-1][j]
			}
		}
	}

	return pascal[rowIndex]
}
func maxProfit(prices []int) int {
	days := len(prices)
	buy := make([]int, days)
	sell := make([]int, days)

	buy[0] = prices[0]

	for i := 1; i < days; i++ {
		buy[i] = min(prices[i], buy[i-1])
	}

	sell[days-1] = prices[days-1]

	for i := days - 2; i >= 0; i-- {
		sell[i] = max(sell[i+1], prices[i])
	}

	ans := 0

	for i := 0; i < days-1; i++ {
		ans = max(ans, sell[i+1]-buy[i])
	}

	return ans
}
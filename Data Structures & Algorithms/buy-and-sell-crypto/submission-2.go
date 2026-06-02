func maxProfit(prices []int) int {
	// DP ver
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0] // min price so far
	best := 0 // best profit so far
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice
		if profit > best {
			best = profit
		}
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
	}
	return best
}

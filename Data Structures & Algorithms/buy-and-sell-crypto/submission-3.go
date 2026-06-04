func maxProfit(prices []int) int {
	// use DP
	if len(prices) == 0 {
		return 0
	}
	lowest, best := prices[0], 0
	for _, price := range prices {
		if lowest > price {
			lowest = price
		}
		diff := price - lowest
		if diff > best {
			best = diff
		}
	}
	return best
}

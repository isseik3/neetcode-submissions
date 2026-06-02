func maxProfit(prices []int) int {
	// we will use two pointers
	left, right := 0, 1
	best := 0
	for right < len(prices) {
		if prices[left] < prices[right] {
			profit := prices[right] - prices[left]
			if profit > best {
				best = profit
			}
		} else {
			left = right
		}
		right++
	}
	return best
}

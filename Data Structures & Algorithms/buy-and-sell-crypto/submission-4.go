func maxProfit(prices []int) int {
	// use two pointers
	left, best := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[left] > prices[i] {
			left = i
		} 
		diff := prices[i] - prices[left]
		if best < diff {
			best = diff
		}
	}
	return best
}

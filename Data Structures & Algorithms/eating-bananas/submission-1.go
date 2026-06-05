func minEatingSpeed(piles []int, h int) int {
	biggest := 0
	for _, pile := range piles {
		if biggest < pile {
			biggest = pile
		}
	}
	hoursNeeded := func(speed int) int {
		hours := 0
		for _, pile := range piles {
			hours += (pile + speed - 1)/speed // ceil(pile/speed)
		}
		return hours
	}
	left, right := 1, biggest
	for left < right {
		mid := left + (right - left)/2
		if hoursNeeded(mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

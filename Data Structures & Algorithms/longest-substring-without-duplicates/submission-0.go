func lengthOfLongestSubstring(s string) int {
	// we will use sliding window
	// last[c] = the index just after we last saw c; used to set the window's left edge
	last := make(map[byte]int)
	left, best := 0, 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		// if this charactor already appeared in the window, move left past its previous position
		if idx, ok := last[c]; ok && idx > left {
			left = idx
		}
		diff := right - left + 1
		if diff > best {
			best = diff
		}
		last[c] = right + 1
	}
	return best
}

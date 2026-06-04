func lengthOfLongestSubstring(s string) int {
	// charSet shows index after its character appears previously
	charSet := make(map[byte]int, len(s))
	left, best := 0, 0
	for i := 0; i < len(s); i++ {
		if idx, ok := charSet[s[i]]; ok && left < idx {
			left = idx
		}
		diff := i - left + 1
		if diff > best {
			best = diff
		}
		charSet[s[i]] = i + 1
	}
	return best

}

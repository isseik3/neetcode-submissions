func characterReplacement(s string, k int) int {
	count := [26]int{}
	left, maxFreq, best := 0, 0, 0
	for right := 0; right < len(s); right++ {
		count[s[right]-'A']++
		if maxFreq < count[s[right]-'A'] {
			maxFreq = count[s[right]-'A']
		}
		// windows invalid if too many chars to replace within budget k
		if right - left + 1 - maxFreq > k {
			count[s[left]-'A']--
			left++
		}
		if best < right - left + 1 {
			best = right - left + 1
		}
	}
	return best
}

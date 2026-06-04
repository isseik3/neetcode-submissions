func characterReplacement(s string, k int) int {
    // sliding window
    counts := [26]int{}
    left, best, maxFreq := 0, 0, 0
    for right := 0; right < len(s); right++ {
        counts[s[right]-'A']++
        if counts[s[right]-'A'] > maxFreq {
            maxFreq = counts[s[right]-'A']
        }
        // invalid windows if too many chars to replace within k budget
        if right - left + 1 - maxFreq > k {
            counts[s[left]-'A']--
            left++
        }
        if best < right - left + 1 {
            best = right - left + 1
        }
    }
    return best
}

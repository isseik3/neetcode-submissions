func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    seen := make(map[rune]int, len(s))
    for _, ss := range s {
        seen[ss]++
    }
    for _, tt := range t {
        seen[tt]--
        if seen[tt] < 0 {
            return false
        }
    }
    return true
}

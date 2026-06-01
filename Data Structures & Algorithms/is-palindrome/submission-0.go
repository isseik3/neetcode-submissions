func isPalindrome(s string) bool {
	runes := make([]rune, 0, len(s))
    for _, ss := range s {
        if isAlphanumeric(ss) {
            runes = append(runes, unicode.ToLower(ss))
        }
    }
    l := len(runes)
    for i := 0; i < l/2; i++ {
        if runes[i] != runes[l-i-1] {
            return false
        }
    }
    return true
}

func isAlphanumeric(s rune) bool {
    return unicode.IsDigit(s) || unicode.IsLetter(s)
}
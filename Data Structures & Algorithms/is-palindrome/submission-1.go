func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        if !isAlphaNum(rune(s[left])) {
            left++
            continue
        }
        if !isAlphaNum(rune(s[right])) {
            right--
            continue
        }
        if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
            return false
        }
        left++
        right--
    }
    return true
}

func isAlphaNum(r rune) bool {
    return unicode.IsDigit(r) || unicode.IsLetter(r)
}

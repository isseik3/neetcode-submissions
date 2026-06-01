func isPalindrome(s string) bool {
    runes := []rune(s)
    left, right := 0, len(runes)-1
    for left < right {
        if !isAlphaNum(runes[left]) {
            left++
            continue
        }
        if !isAlphaNum(runes[right]) {
            right--
            continue
        }
        if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
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
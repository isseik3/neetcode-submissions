func isValid(s string) bool {
    pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
    stack := []byte{}

    for i := 0; i < len(s); i++ {
        c := s[i]
        if open, ok := pairs[c]; ok {
            if len(stack) == 0 || open != stack[len(stack)-1] {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, c)
        }
    }
    return len(stack) == 0

}
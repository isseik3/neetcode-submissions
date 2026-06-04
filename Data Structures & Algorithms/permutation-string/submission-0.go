func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    s1Count, s2Count := [26]int{}, [26]int{}
    for i := 0; i < len(s1); i++ {
        s1Count[s1[i]-'a']++
        s2Count[s2[i]-'a']++
    }

    matches := 0
    for i := 0; i < 26; i++ {
        if s1Count[i] == s2Count[i] {
            matches++
        }
    }

    left := 0
    for right := len(s1); right < len(s2); right++ {
        if matches == 26 {
            return true
        }
        idx := s2[right] - 'a'
        s2Count[idx]++
        if s2Count[idx] == s1Count[idx] {
            matches++
        } else if s1Count[idx]+1 == s2Count[idx] {
            matches--
        }

        idx = s2[left] - 'a'
        s2Count[idx]--
        if s2Count[idx] == s1Count[idx] {
            matches++
        } else if s1Count[idx]-1 == s2Count[idx] {
            matches--
        }
        left++
    }
    return matches == 26
}
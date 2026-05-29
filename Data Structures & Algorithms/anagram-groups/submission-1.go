func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)
    
    for _, str := range strs {
        var count [26]int
        for _, s := range str {
            count[s-'a']++
        }
        groups[count] = append(groups[count], str)
    }
    ans := make([][]string, 0, len(groups))
    for _, group := range groups {
        ans = append(ans, group)
    }
    return ans
}

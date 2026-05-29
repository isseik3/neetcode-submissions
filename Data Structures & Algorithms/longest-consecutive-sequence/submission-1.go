func longestConsecutive(nums []int) int {
    numSets := make(map[int]struct{}, len(nums))

    for _, num := range nums {
        numSets[num] = struct{}{}
    }

    longest := 0
    for n := range numSets {
        if _, ok := numSets[n-1]; !ok { // start of a sequence
            length := 1
            for {
                if _, ok := numSets[n+length]; !ok {
                    break
                }
                length++
            }
            if length > longest {
                longest = length
            }
        }
    }
    return longest
}

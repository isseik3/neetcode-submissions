func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		numSet[n] = struct{}{}
	}

	longest := 0
	for n := range numSet {
		if _, ok := numSet[n-1]; !ok { // start of a sequence
			length := 1
			for {
				if _, ok := numSet[n+length]; !ok {
					break
				}
				length++
			}
			if longest < length {
				longest = length
			}
		}
	}
	return longest
}

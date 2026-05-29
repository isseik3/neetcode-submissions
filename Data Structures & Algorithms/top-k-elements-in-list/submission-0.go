func topKFrequent(nums []int, k int) []int {
    x := make(map[int]int, len(nums))
    for _, num := range nums {
        x[num]++
    }
    
    keys := make([]int, 0, len(x))
    for xx := range x {
        keys = append(keys, xx)
    }

    sort.Slice(keys, func(i, j int) bool {
        return x[keys[i]] > x[keys[j]]
    })
    return keys[:k]
}

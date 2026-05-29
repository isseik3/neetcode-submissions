func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int, len(nums))
    for _, num := range nums {
        count[num]++
    }

    // buckets[i] = numbers that appear exactly i times
    buckets := make([][]int, len(nums)+1)
    for num, c := range count {
        buckets[c] = append(buckets[c], num)
    }

    // Walk from highest to lowerest, collect k elements
    result := make([]int, 0, k)
    for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
        result = append(result, buckets[i]...)
    }

    return result[:k]
}

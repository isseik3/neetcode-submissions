func twoSum(nums []int, target int) []int {
    x := make(map[int]int, len(nums))
    for i, n := range nums {
        complement := target - n
        if j, ok := x[complement]; ok {
            return []int{j, i}
        }
        x[n] = i
    }
    return []int{}
}
